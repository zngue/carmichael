package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_km"
)

/*
*@Author Administrator
*@Date 9/4/2021 11:48
*@desc
 */
func ZngKmRouter(group *gin.RouterGroup) {
	ZngKmGroup := group.Group("zngKm")
	{
		ZngKmGroup.GET("list", zng_km.List)
		ZngKmGroup.GET("detail", zng_km.Detail)
		ZngKmGroup.POST("edit", zng_km.Edit)
		ZngKmGroup.POST("delete", zng_km.Delete)
		ZngKmGroup.POST("add", zng_km.Add)
	}
}
