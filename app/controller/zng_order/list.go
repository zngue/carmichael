package zng_order

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)
func  List(ctx *gin.Context) {
	var req request.ZngOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	res, err := service.NewZngOrderService().List(&req)
	response.HttpSuccessWithError(ctx,err,res)
	return
}
