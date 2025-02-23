package routers

import (
	controller "github.com/HotfixHoudini/reader-bot/pkg/controllers/greeting"
	"github.com/gin-gonic/gin"
)

func Greeting(router *gin.Engine) *gin.Engine {
	greetController := controller.Greet

	{
		router.GET("/", greetController)
	}

	return router
}
