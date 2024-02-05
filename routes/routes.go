package routes

import (
	"go-k8s-game/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	vuUser := router.Group("user")
	{
		vuUser.GET("/health", handler.GetHealth)

		v1 := vuUser.Group("v1")
		{

			v1.GET("sum", handler.GetSum)
		}
	}
}
