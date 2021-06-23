package router

import "github.com/gin-gonic/gin"

func Router(group *gin.RouterGroup) {
	ZngCateRouter(group)
	ZngOrderRouter(group)
	ZngShopRouter(group)
	ZngShopCarmichaelRouter(group)
}
