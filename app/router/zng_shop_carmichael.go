package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_shop_carmichael"
)

// ZngShopCarmichaelRouter /*
func ZngShopCarmichaelRouter(group *gin.RouterGroup) {
	ZngShopCarmichaelGroup := group.Group("zngShopCarmichael")
	{
		ZngShopCarmichaelGroup.GET("list", zng_shop_carmichael.List)
		ZngShopCarmichaelGroup.GET("detail", zng_shop_carmichael.Detail)
		ZngShopCarmichaelGroup.POST("edit", zng_shop_carmichael.Edit)
		ZngShopCarmichaelGroup.POST("delete", zng_shop_carmichael.Delete)
		ZngShopCarmichaelGroup.POST("add", zng_shop_carmichael.Add)
	}
}
