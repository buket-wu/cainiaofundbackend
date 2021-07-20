package user

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	userR := r.Group("/user")

	userR.POST("/addUser", AddUser)
	userR.GET("/getUser", GetUser)
	userR.POST("/editUser", EditUser)
	userR.GET("/delUser", DelUser)
	userR.POST("/addMyFund", AddMyFund)
	userR.GET("/getMyFund", GetMyFund)
	userR.GET("/getMyFundList", GetMyFundList)
	userR.POST("/editMyFund", EditMyFund)
	userR.GET("/delMyFund", DelMyFund)
}
