package zng_user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)

func List(ctx *gin.Context) {
	var req request.ZngUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	res, err := service.NewZngUserService().List(&req)
	response.HttpSuccessWithError(ctx, err, res)
	return
}
