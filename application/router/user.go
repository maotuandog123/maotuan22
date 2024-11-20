package router

import (
	"github.com/gin-gonic/gin"
	"urils/application/api"
)

func InitRouter(Router *gin.RouterGroup) {
	userRoute := Router.Group("user")
	{
		userRoute.POST("authenticate", api.UserAuthenticate)
	}
}
