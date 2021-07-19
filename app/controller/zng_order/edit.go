package zng_order

import (
	"encoding/json"
	"errors"

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
	var req request.ZngOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	postForm := ctx.PostForm("updateData")
	if len(postForm) == 0 {
		response.HttpParameterError(ctx, errors.New("update data 参数错误"))
		return
	}
	var ma map[string]interface{}
	if err := json.Unmarshal([]byte(postForm), &ma); err != nil {
		response.HttpParameterError(ctx, err)
		return
	}
	req.Data = ma
	req.ReturnType = 3
	err := service.NewZngOrderService().Edit(&req)
	response.HttpSuccessWithError(ctx, err, nil)
}
