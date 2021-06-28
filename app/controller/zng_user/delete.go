package zng_user

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
	var req request.ZngUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	err := service.NewZngUserService().Delete(&req)
	response.HttpSuccessWithError(ctx, err, nil)
}
