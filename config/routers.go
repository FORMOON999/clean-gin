package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	// not found
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "Not found")
	})

	return r
}
