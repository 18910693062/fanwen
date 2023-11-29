package router

import (
	"github.com/gin-gonic/gin"
	"go-edu/config"
	"go-edu/controller"
)

func Setup() *gin.Engine {
	gin.SetMode(config.Conf.Mode)
	r := gin.New()

	r.GET("/hello", controller.HelloHandler)
	return r
}
