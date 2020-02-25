package router

import (
	"dino-infra/handler"

	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/ping", handler.PingPongHandler)
}
