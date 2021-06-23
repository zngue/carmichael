package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_cate"
)

/*
*@Author Administrator
*@Date 9/4/2021 11:48
*@desc
 */
func ZngCateRouter(group *gin.RouterGroup)  {
	ZngCateGroup := group.Group("zngCate")
	{
		ZngCateGroup.GET("list",zng_cate.List)
		ZngCateGroup.GET("detail",zng_cate.Detail)
		ZngCateGroup.POST("edit",zng_cate.Edit)
		ZngCateGroup.POST("delete",zng_cate.Delete)
		ZngCateGroup.POST("add",zng_cate.Add)
	}
}
