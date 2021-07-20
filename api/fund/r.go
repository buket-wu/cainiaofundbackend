package fund

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	fundR := r.Group("/fund")

	fundR.POST("/addFund", AddFund)
	fundR.GET("/getFundList", GetFundList)
}
