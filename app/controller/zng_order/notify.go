package zng_order

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/zngue/carmichael/app/httplib"

	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"

	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/go_helper/pkg"
	"golang.org/x/sync/errgroup"
)

type Response struct {
	MchIdNo        string `json:"mch_id_no"`
	Money          int    `json:"money"`
	Openid         string `json:"openid"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	SuccessTime    string `json:"success_time"`
}

func sendTemplate(maps map[string]string) {

	sendMessageUrl := viper.GetString("payment.sendMessageUrl")
	if len(sendMessageUrl) == 0 {
		fmt.Println("sendMessageUrl null")
		return
	}
	httpRequest := httplib.Get(sendMessageUrl)
	if maps != nil {
		for key, val := range maps {
			httpRequest.Param(key, val)
		}
	}
	httpRequest.Response()
}

func Notify(ctx *gin.Context) {
	var rep Response
	if err := ctx.BindJSON(&rep); err != nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "参数接受失败",
		})
		return
	}
	begin := pkg.MysqlConn.Begin()
	var shopReq request.ZngOrderRequest
	shopReq.OrderNum = rep.MchIdNo
	detail, err := service.NewZngOrderService().Detail(&shopReq)
	if err != nil || detail == nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	}
	var kmOne model.ZngKm
	var shopOne model.ZngShop

	var shopwg errgroup.Group
	shopwg.Go(func() error { //获取卡密
		return pkg.MysqlConn.Model(&model.ZngKm{}).Where("status = ?", 0).Where(" shop_id =  ? ", detail.ShopId).First(&kmOne).Error
	})
	shopwg.Go(func() error { // 获取商品信息
		return pkg.MysqlConn.Model(&model.ZngShop{}).Where("id = ?", detail.ShopId).First(&shopOne).Error
	})
	if err := shopwg.Wait(); err != nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	}
	if &kmOne == nil || &shopOne == nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	}
	defer begin.Rollback()
	var wg errgroup.Group
	wg.Go(func() error { //更新订单信息
		return begin.Model(&model.ZngOrder{}).Where("order_num = ?", rep.MchIdNo).Updates(map[string]interface{}{
			"pay_status":  1,
			"update_time": time.Now().Unix(),
			"openid":      rep.Openid,
		}).Error

	})
	wg.Go(func() error {
		return begin.Model(&model.ZngShop{}).Where("id = ?", detail.ShopId).Updates(map[string]interface{}{
			"sales_num": shopOne.SalesNum + 1,
		}).Error
	})
	if kmOne.KmType == 2 { // idea 账号密码不更新数据库
		wg.Go(func() error {
			return begin.Model(&model.ZngUser{}).Where("openid = ?", rep.Openid).Updates(map[string]interface{}{
				"account":      kmOne.Account,
				"password":     kmOne.Password,
				"ext_account":  kmOne.ExtAccount,
				"ext_password": kmOne.ExtPassword,
				"code_content": kmOne.CodeContent,
				"expire_time":  time.Now().Unix() + kmOne.ExpireTime,
			}).Error
		})
	} else {
		wg.Go(func() error {
			return begin.Model(&model.ZngUser{}).Where("openid = ?", rep.Openid).Updates(map[string]interface{}{
				"account":     kmOne.Account,
				"password":    kmOne.Password,
				"expire_time": time.Now().Unix() + kmOne.ExpireTime,
			}).Error
		})
		wg.Go(func() error {
			return begin.Model(&model.ZngKm{}).Where("id = ?", kmOne.Id).Updates(map[string]interface{}{
				"Status": 1,
			}).Error
		})
	}
	err = wg.Wait()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	}
	errdb := begin.Commit().Error
	if errdb != nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	}

	sendTemplate(map[string]string{
		"account":  kmOne.Account,
		"password": kmOne.Password,
		"openid":   rep.Openid,
	})
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "支付成功",
	})
	return

}
