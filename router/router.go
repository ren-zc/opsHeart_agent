package router

import (
	"github.com/gin-gonic/gin"
	hv1 "opsHeart_agent/handlers/v1"
	"opsHeart_agent/router/middleware"
)

var RegRouter *gin.Engine
var R *gin.Engine

func init() {
	RegRouter = gin.Default()
	RegV1 := RegRouter.Group("/v1")
	{
		RegV1.POST("/start-up", hv1.HandleStartUp)
	}

	R = gin.Default()
	v1 := R.Group("/v1")
	v1.Use(middleware.TokenChecker())
	{
		v1.POST("/net-check", nil)
		v1.POST("/cmd", nil)
		v1.POST("/script/run", nil)
	}
}
