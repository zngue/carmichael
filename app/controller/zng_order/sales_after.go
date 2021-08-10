package zng_order

import (
	"fmt"
	"strings"
	"time"

	"github.com/zngue/carmichael/app/httplib"

	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/go_helper/pkg"
)

func InArrayString(arr []string, s string) bool {

	if len(arr) == 0 {
		return false
	}
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false

}
func SaleAfter(ctx *gin.Context) {

	var list []model.ZngOrder
	err := pkg.MysqlConn.Model(&model.ZngOrder{}).Where("openid !=  ''").Where("pay_status = ?", 1).Order("id desc").Find(&list).Error
	fmt.Println(err)
	Begin := pkg.MysqlConn.Begin()
	defer Begin.Rollback()
	var openidArr []string
	nowTime := time.Now().Unix()
	for _, order := range list {
		var one model.ZngKm
		if InArrayString(openidArr, order.OpenID) {
			continue
		}
		err := pkg.MysqlConn.Model(&model.ZngKm{}).Where("shop_id = ?", order.ShopId).Find(&one).Error
		fmt.Println(err)

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
				openidArr = append(openidArr, order.OpenID)
			}
		}
	}
	fmt.Println(strings.Join(openidArr, ","))

	SalesSendTemplate(map[string]string{
		"openid": strings.Join(openidArr, ","),
	})
	/*
		Begin.Commit()*/
}
func SalesSendTemplate(maps map[string]string) {
	httpRequest := httplib.Get("http://127.0.0.1:6060/pay/message/salesAfter")
	if maps != nil {
		for key, val := range maps {
			httpRequest.Param(key, val)
		}
	}
	httpRequest.Response()
}
