package main

import (
	"time"

	"github.com/auth-core/app/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

var runMode string
var theTime time.Time

func initRoutes(debug bool, mode string) {
	//serverMode := config.MustGetString("server.mode")

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/", Ping)
	v1 := router.Group("/v1")
	v1.POST("/auth/token", controllers.GetToken)

	v1Auth := router.Group("/v1")
	v1Auth.Use(controllers.Authenticate)
	// v1Auth.Use(controllers.Logging)
	{
		v1Auth.GET("/location/getcity/:coordinate/:radius", controllers.GetCity)
	}
}
