package zng_order

import (
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/go_helper/pkg"

	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)

//售后处理
func Sales(ctx *gin.Context) {
	var req request.ZngOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	orderInfo, errs := service.NewZngOrderService().Detail(&req) //获取条订单消息
	if orderInfo == nil || errs != nil {
		response.HttpParameterError(ctx, errors.New("参数错误"))
	}
	//获取卡密信息
	var kmOne model.ZngKm
	err := pkg.MysqlConn.Model(&model.ZngKm{}).Where("status = ?", 0).Where(" shop_id =  ? ", orderInfo.ShopId).First(&kmOne).Error
	if err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	var shopwg errgroup.Group
	//更新卡密 或者更新数据信息
	begin := pkg.MysqlConn.Begin()
	defer begin.Rollback()
	shopwg.Go(func() error {
		return begin.Model(&model.ZngOrder{}).Where("order_num = ?", req.OrderNum).Updates(map[string]interface{}{
			"is_sell_after": 2,
		}).Error
	})
	shopwg.Go(func() error {
		return begin.Model(&model.ZngUser{}).Where("openid = ?", orderInfo.OpenID).Updates(map[string]interface{}{
			"account":      kmOne.Account,
			"password":     kmOne.Password,
			"ext_account":  kmOne.ExtAccount,
			"ext_password": kmOne.ExtPassword,
		}).Error
	})
	shoerr := shopwg.Wait()
	if shoerr != nil {
		response.HttpFailWithErr(ctx, shoerr)
		return
	}
	shoerr = begin.Commit().Error
	if shoerr != nil {
		response.HttpFailWithErr(ctx, shoerr)
		return
	}
	response.HttpOkWithMessage(ctx, "操作成功")
	return

}
