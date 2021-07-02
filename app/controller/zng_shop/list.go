package zng_shop

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)

func List(ctx *gin.Context) {
	var req request.ZngShopRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	if req.Status == 0 {
		req.Status = 1
	}
	res, err := service.NewZngShopService().List(&req)
	response.HttpSuccessWithError(ctx, err, res)
	return
}
