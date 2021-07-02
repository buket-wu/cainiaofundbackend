package api

import (
	"cainiaofundbackend/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var server *gin.Engine

func init() {
	server = gin.New()
	server.Use(middleware.LogMiddleware())
	baseGroup := server.Group("/api")
	baseGroup.GET("/", health)
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "service is up")
}

func GetServer() *gin.Engine {
	return server
}
