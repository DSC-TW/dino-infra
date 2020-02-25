package handler

import "github.com/gin-gonic/gin"

/*
PingPongHandler return pong when user hit ping
*/
func PingPongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
