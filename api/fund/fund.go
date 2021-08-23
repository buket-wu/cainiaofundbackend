package fund

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/utils"
	"cainiaofundbackend/xiong"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddFund(c *gin.Context) {
	req := AddFundReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "code require")
		return
	}

	xReq := &xiong.GetFundReq{
		Code: req.Code,
	}

	fundList, err := xiong.GetFund(xReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "get fund fail")
		return
	}

	insertMany := make([]interface{}, 0)
	for _, fund := range fundList {
		insert := &db.Fund{
			ID:         primitive.NewObjectID(),
			Code:       fund.Code,
			Name:       fund.Name,
			Status:     db.FundStatusOn,
			Createtime: utils.Now(),
			Updatetime: utils.Now(),
		}
		insertMany = append(insertMany, insert)
	}

	res, err := db.FundCol.InsertMany(c, insertMany)
	if err != nil {
		logrus.Errorf("err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "insert fail")
		return
	}

	c.JSON(http.StatusOK, res.InsertedIDs)
	return
}

func GetFundList(c *gin.Context) {
	rsp := []db.Fund{}
	cur, err := db.FundCol.Find(c, bson.D{{}})
	if err != nil {
		logrus.Errorf("err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "query fail")
		return
	}

	err = cur.All(c, &rsp)
	if err != nil {
		logrus.Errorf("err:%v", err)
	}

	// Close the cursor once finished
	_ = cur.Close(c)

	c.JSON(http.StatusOK, rsp)
	return
}
