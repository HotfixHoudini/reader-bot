package controllers

import (
	"net/http"

	services "github.com/HotfixHoudini/reader-bot/services/integration"
	"github.com/gin-gonic/gin"
)

type IntegrationJSONController struct {
	Service *services.IntegrationJsonService
}

func (inc *IntegrationJSONController) GetIntegrationJSON(c *gin.Context) {

	data := inc.Service.GetIntegrationsJSON(c)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
