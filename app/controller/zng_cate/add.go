package zng_cate

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/carmichael/app/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func Add(ctx *gin.Context) {
	var req request.ZngCateRequest
	var data model.ZngCate
	if err := ctx.ShouldBind(&data); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	req.Data=&data
	err := service.NewZngCateService().Add(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
