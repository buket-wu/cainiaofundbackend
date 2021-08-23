package user

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddUser(c *gin.Context) {
	req := AddUserReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "code require")
		return
	}

	insert := &db.User{
		ID:         primitive.NewObjectID(),
		Username:   req.Username,
		Openid:     "",
		Createtime: utils.Now(),
		Updatetime: utils.Now(),
	}

	res, err := db.UserCol.InsertOne(c, insert)
	if err != nil {
		logrus.Errorf("err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "insert fail")
		return
	}

	c.JSON(http.StatusOK, res)
	return

}

func GetUser(c *gin.Context) {

}

func EditUser(c *gin.Context) {

}

func DelUser(c *gin.Context) {

}

func AddMyFund(c *gin.Context) {

}

func GetMyFundList(c *gin.Context) {

}

func GetMyFund(c *gin.Context) {

}

func EditMyFund(c *gin.Context) {

}

func DelMyFund(c *gin.Context) {

}
