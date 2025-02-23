package routers

import (
	controllers "github.com/HotfixHoudini/reader-bot/pkg/controllers/integration"
	"github.com/gin-gonic/gin"
)

func Integration(router *gin.Engine) *gin.Engine {
	IC := controllers.IntegrationJSONController{}

	{
		router.GET("/integration", IC.GetIntegrationJSON)
	}

	return router
}
