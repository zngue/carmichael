package zng_order

import (
	"time"

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
	if shopOne.CateId == 1 { // idea 账号密码不更新数据库
		wg.Go(func() error {
			return begin.Model(&model.ZngUser{}).Where("openid = ?", rep.Openid).Updates(map[string]interface{}{
				"account":      kmOne.Account,
				"password":     kmOne.Password,
				"ext_account":  kmOne.ExtAccount,
				"ext_password": kmOne.ExtPassword,
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

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "支付成功",
	})
	return

}
