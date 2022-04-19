package router

import (
	"router/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Pong)
	r.GET("/nico", controller.Nico)
	r.POST("/login", controller.Login)
	return r
}
