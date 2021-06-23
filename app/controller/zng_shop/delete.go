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
func Delete(ctx *gin.Context) {
	var req request.ZngShopRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	err := service.NewZngShopService().Delete(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
