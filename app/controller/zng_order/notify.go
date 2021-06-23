package zng_order

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/go_helper/pkg"
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
	err := pkg.MysqlConn.Model(&model.ZngOrder{}).Where("order_num = ?", rep.MchIdNo).Updates(map[string]interface{}{
		"pay_status":  1,
		"update_time": time.Now().Unix(),
		"openid":      rep.Openid,
	}).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    100,
			"message": "更新失败",
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "支付成功",
		})
		return
	}

}
