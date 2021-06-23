package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_shop"
)

/*
*@Author Administrator
*@Date 9/4/2021 11:48
*@desc
 */
func ZngShopRouter(group *gin.RouterGroup)  {
	ZngShopGroup := group.Group("zngShop")
	{
		ZngShopGroup.GET("list",zng_shop.List)
		ZngShopGroup.GET("detail",zng_shop.Detail)
		ZngShopGroup.POST("edit",zng_shop.Edit)
		ZngShopGroup.POST("delete",zng_shop.Delete)
		ZngShopGroup.POST("add",zng_shop.Add)
	}
}
