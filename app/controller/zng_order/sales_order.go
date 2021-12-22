package zng_order

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/response"
	"gorm.io/gorm"
	"time"
)

func SalesOrder(ctx *gin.Context) {
	orderNum := ctx.DefaultQuery("orderNum", "")
	if len(orderNum) == 0 {
		response.HttpFailWithMessage(ctx, "请输入订单号")
		return
	}
	var order model.ZngOrder
	err := pkg.MysqlConn.Model(&model.ZngOrder{}).Where("order_num  =  ?", orderNum).Where("pay_status = ?", 1).First(&order).Error
	if err == gorm.ErrRecordNotFound {
		response.HttpFailWithMessage(ctx, "没有可售后的订单")
		return
	}
	if err != nil {
		response.HttpFailWithMessage(ctx, "订单编号错误")
		return
	}
	Begin := pkg.MysqlConn.Begin()
	defer Begin.Rollback()
	nowTime := time.Now().Unix()
	var one model.ZngKm
	errs := pkg.MysqlConn.Model(&model.ZngKm{}).Where("shop_id = ?", order.ShopId).Find(&one).Error
	if errs != nil {
		response.HttpFailWithMessage(ctx, "参数错误")
		return
	}
	if nowTime < order.AddTime+one.ExpireTime {
		var user model.ZngUser
		pkg.MysqlConn.Model(&user).Where(" openid = ? ", order.OpenID).First(&user)
		updateUser := make(map[string]interface{})
		if &user != nil {
			updateUser["account"] = one.Account
			updateUser["ext_account"] = one.ExtAccount
			updateUser["password"] = one.Password
			updateUser["ext_password"] = one.ExtPassword
			if user.ExpireTime == 0 || user.ExpireTime != order.AddTime+one.ExpireTime {
				updateUser["expire_time"] = order.AddTime + one.ExpireTime
			}
			Begin.Model(&model.ZngUser{}).Where(" openid = ? ", order.OpenID).Updates(updateUser)
		}
	} else {
		response.HttpFailWithMessage(ctx, "订单已过期")
		return
	}
	Begin.Commit()
	SalesSendTemplate(map[string]string{
		"openid": order.OpenID,
	})
	response.HttpOk(ctx)
	return
}
