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
	userService := service.NewZngUserService()
	req.Openid = data.Openid
	detail, err2 := userService.Detail(&req)

	if err2 != nil && err2 != gorm.ErrRecordNotFound { //有错误但是不是查询数据不存在
		response.HttpParameterError(ctx, err2)
		return
	}
	req.Data = &model.ZngUser{}
	req.Data = &data
	num := req.GetDB().RowsAffected

	if num > 0 {
		if len(data.Unionid) > 0 {
			req.Data = map[string]interface{}{
				"unionid": data.Unionid,
			}
			req.ID = detail.Id
			if err3 := userService.Edit(&req); err3 != nil {
				response.HttpParameterError(ctx, err3)
				return
			}
		} else {
			response.HttpOkWithMessage(ctx, "用户已存在", detail)
			return
		}
	} else {
		err := userService.Add(&req)
		if err != nil {
			response.HttpFailWithMessage(ctx, err.Error())
			return
		}
		response.HttpSuccessWithError(ctx, err, data)
	}
	if num > 0 {
		response.HttpOkWithMessage(ctx, "用户已存在", detail)
		return
	}

	return
}
