package user

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/extend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AddUser(c *gin.Context) {
	req := AddUserReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "code require")
		return
	}

	insert := &db.User{
		Username:   req.Username,
		Openid:     uuid.New().String(),
		Createtime: utils.Now(),
		Updatetime: utils.Now(),
	}

	res, err := db.GetUserCol().InsertOne(c, insert)
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
