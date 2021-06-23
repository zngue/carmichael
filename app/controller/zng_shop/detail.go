package zng_shop

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func  Detail(ctx *gin.Context) {
	var req request.ZngShopRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	res,err := service.NewZngShopService().Detail(&req)
	response.HttpSuccessWithError(ctx,err,res)
	return
}
