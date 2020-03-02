package main

import (
	"dino-infra/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var httpServer *gin.Engine

	httpServer = gin.Default()

	router.Register(httpServer)

	serverAddr := "localhost:8080"
	httpServer.Run(serverAddr)
}
