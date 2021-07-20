package job

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	userR := r.Group("/job")

	userR.GET("/sync", Sync)
}
