package zng_order

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/zngue/carmichael/app/httplib"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/xinliangnote/go-util/md5"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)

func randomMath(end, start int64) int64 {
	randoms := rand.Int63n(end - start)
	return randoms + start
}
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
func OrderNum() string {
	return "km" + time.Now().Format("20060102150405") + cast.ToString(randomMath(999999, 100000)) + cast.ToString(randomMath(999999, 100000))
}

type OrderRequest struct {
	ShopID int    `json:"shop_id" form:"shopId"`
	UserID int64  `json:"user_id" form:"userId"`
	OpenID string `form:"openid"`
}

// Add /*
func Add(ctx *gin.Context) {
	/*
		    'app_id'=>$appid,
			'sign'=>md5($appkey.$appid.$time.$orderStr),
			"signTime"=>$time,
			"signStr"=>$orderStr,
			"mch_id_no"=>$orderNo,
			"money"=>999,
			"order_desc"=>"idea年账号购买",
			"notify_url"=>"http://bbs.zngue.com/api/pay/orderNotify",
			"return_url"=>"http://bbs.zngue.com/set",
	*/
	var req request.ZngOrderRequest
	var data OrderRequest
	if err := ctx.ShouldBind(&data); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}

	shopRequest := request.ZngShopRequest{
		ID: data.ShopID,
	}

	shop, err3 := service.NewZngShopService().Detail(&shopRequest)
	if err3 != nil {
		response.HttpFailWithMessage(ctx, err3.Error())
		return
	}
	if shop.UserLimit > 0 && data.OpenID != "oPv356jetWWZPjVY31e7Eiuy3kZQ" {
		countNum, err := service.NewZngOrderService().Count(data.OpenID)
		if err != nil {
			response.HttpFailWithMessage(ctx, err.Error())
			return
		}
		if countNum >= shop.UserLimit {
			response.HttpFailWithMessage(ctx, fmt.Sprintf("体验卡限购%d次", shop.UserLimit))
			return
		}
	}
	var orderNo = OrderNum()
	order := model.ZngOrder{
		OrderNum:        orderNo,
		ShopId:          shop.Id,
		UserId:          data.UserID,
		ShopTotailPrice: float64(shop.Money),
		ShopNum:         1,
		ShopPrice:       float64(shop.Money),
		AddTime:         time.Now().Unix(),
		OrderStatus:     1,
		PayStatus:       0,
		ShopTitle:       shop.Title,
	}
	appid := viper.GetString("payment.appid")
	appkey := viper.GetString("payment.appkey")
	signTime := time.Now().Unix()
	signStr := RandString(15)
	var crData = map[string]interface{}{
		"app_id":     appid,
		"sign":       md5.MD5(appkey + appid + cast.ToString(signTime) + signStr),
		"signTime":   signTime,
		"signStr":    signStr,
		"mch_id_no":  orderNo,
		"money":      order.ShopTotailPrice,
		"order_desc": order.ShopTitle,
		"notify_url": viper.GetString("payment.notify_url"),
		"return_url": viper.GetString("payment.return_url"),
	}
	url := viper.GetString("payment.createUrl")
	httpRequest, err2 := httplib.Post(url).JSONBody(crData)
	if err2 != nil {
		response.HttpFailWithMessage(ctx, err2.Error())
		return
	}
	bytes, err2 := httpRequest.Bytes()
	if err2 != nil {
		response.HttpFailWithMessage(ctx, err2.Error())
		return
	}
	var OrderRes resp
	if err := json.Unmarshal(bytes, &OrderRes); err != nil {
		response.HttpFailWithMessage(ctx, err.Error())
		return
	}

	req.Data = &order
	err := service.NewZngOrderService().Add(&req)
	if err == nil && (OrderRes.Code == 200 || OrderRes.Code == 203) {
		paymentUrl := viper.GetString("payment.paymentUrl") + "?mchIdNo=" + orderNo
		response.HttpSuccessWithError(ctx, err, map[string]interface{}{
			"paymentUrl": paymentUrl,
			"orderNo":    orderNo,
		})
		return
	} else {
		response.HttpFailWithMessage(ctx, "操作失败")
		return
	}

}

type resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
