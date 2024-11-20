package initialize

import (
	"github.com/gin-gonic/gin"
	"urils/application/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/api")
	router.InitRouter(ApiGroup)
	return Router
}
