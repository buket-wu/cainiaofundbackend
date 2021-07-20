package api

import (
	"cainiaofundbackend/api/fund"
	"cainiaofundbackend/api/job"
	"cainiaofundbackend/api/user"
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
	fund.RegisterRouter(baseGroup)
	user.RegisterRouter(baseGroup)

	internal := server.Group("/internal")
	job.RegisterRouter(internal)
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "service is up")
}

func GetServer() *gin.Engine {
	return server
}
