package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_user"
)

// ZngUserRouter /*
func ZngUserRouter(group *gin.RouterGroup) {
	ZngUserGroup := group.Group("zngUser")
	{
		ZngUserGroup.GET("list", zng_user.List)
		ZngUserGroup.GET("detail", zng_user.Detail)
		ZngUserGroup.POST("edit", zng_user.Edit)
		ZngUserGroup.POST("delete", zng_user.Delete)
		ZngUserGroup.POST("add", zng_user.Add)
	}
}
