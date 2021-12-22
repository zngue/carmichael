package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/carmichael/app/controller/zng_order"
)

/*
*@Author Administrator
*@Date 9/4/2021 11:48
*@desc
 */
func ZngOrderRouter(group *gin.RouterGroup) {
	ZngOrderGroup := group.Group("zngOrder")
	{
		ZngOrderGroup.GET("list", zng_order.List)
		ZngOrderGroup.GET("detail", zng_order.Detail)
		ZngOrderGroup.POST("edit", zng_order.Edit)
		ZngOrderGroup.POST("delete", zng_order.Delete)
		ZngOrderGroup.POST("add", zng_order.Add)
		ZngOrderGroup.POST("notify", zng_order.Notify)
		ZngOrderGroup.POST("sales", zng_order.Sales)
		ZngOrderGroup.GET("saleAfter", zng_order.SaleAfter)
		ZngOrderGroup.GET("salesOrder", zng_order.SalesOrder)
	}
}
