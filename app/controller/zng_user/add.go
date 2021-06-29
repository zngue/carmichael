package zng_user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
	"gorm.io/gorm"
)

// Add /*
func Add(ctx *gin.Context) {
	var req request.ZngUserRequest
	var data model.ZngUser
	if err := ctx.ShouldBind(&data); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	req.Data = &data
	userService := service.NewZngUserService()
	req.Openid = data.Openid
	detail, err2 := userService.Detail(&req)
	if err2 != nil && err2 != gorm.ErrRecordNotFound { //有错误但是不是查询数据不存在
		response.HttpParameterError(ctx, err2)
		return
	}
	if detail != nil {
		response.HttpOkWithMessage(ctx, "用户已存在", detail)
		return
	}
	err := userService.Add(&req)
	if err != nil {
		response.HttpFailWithMessage(ctx, err.Error())
		return
	}
	response.HttpSuccessWithError(ctx, err, data)
	return
}
