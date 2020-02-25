package main

import (
	"dino-infra/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var httpServer *gin.Engine

	httpServer = gin.Default()

	router.Register(httpServer)

	serverAddr := "0.0.0.0:8080"
	httpServer.Run(serverAddr)
}
