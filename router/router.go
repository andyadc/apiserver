package router

import (
	"apiserver/handler/sd"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(e *gin.Engine, hf ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	e.Use(gin.Recovery())
	e.Use(middleware.NoCache)
	e.Use(middleware.Options)
	e.Use(middleware.Secure)
	e.Use(hf...)

	// 404 Handler.
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	// The health check handlers
	svcd := e.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return e
}
