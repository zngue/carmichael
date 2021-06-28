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
func Edit(ctx *gin.Context) {
	var req request.ZngUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	if err := ctx.Request.ParseForm(); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	postForm := ctx.Request.PostForm
	data := make(map[string]interface{})
	for key, val := range postForm {
		data[key] = val
	}
	req.Data = data
	err := service.NewZngUserService().Edit(&req)
	response.HttpSuccessWithError(ctx, err, nil)
}
