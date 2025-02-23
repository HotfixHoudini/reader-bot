package main

import (
	"github.com/HotfixHoudini/reader-bot/pkg/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routing.Route()

	r.Run("0.0.0.0:8080")
}
