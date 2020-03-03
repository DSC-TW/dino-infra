package router

import (
	"dino-infra/handler"
	"dino-infra/middleware"

	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine) {
	router.Use(middleware.NewCORSMiddleware())

	api := router.Group("/api")

	api.GET("/ping", handler.PingPongHandler)
	api.POST("/user", handler.PushUser)
	api.GET("/user", handler.GetUser)
	api.POST("/score", handler.UpdateScore)
	api.GET("/score", handler.GetScore)
	api.GET("/rank", handler.GetRank)
}
