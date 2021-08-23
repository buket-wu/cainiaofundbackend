package fund

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/utils"
	"cainiaofundbackend/xiong"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func AddFund(c *gin.Context) {
	req := AddFundReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	resp := AddFundResp{}

	codeArr := strings.Split(req.Codes, ",")
	codeMap := make(map[string]int)
	for _, code := range codeArr {
		codeMap[code] = 1
	}

	newCodeList := make([]string, 0)
	cur, err := db.GetFundCol().Find(c, bson.M{"code": bson.M{"$in": codeArr}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "get local fund fail")
		return
	}
	for cur.Next(c) {
		var fund db.Fund
		err := cur.Decode(&fund)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "decode fund fail")
			return
		}

		if _, ok := codeMap[fund.Code]; !ok {
			newCodeList = append(newCodeList)
		}

	}

	if len(newCodeList) == 0 {
		c.JSON(http.StatusOK, resp)
		return
	}

	xReq := &xiong.GetFundReq{
		Code: strings.Join(newCodeList, ","),
	}

	fundList, err := xiong.GetFund(xReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "get fund fail")
		return
	}

	insertMany := make([]interface{}, 0)
	for _, fund := range fundList {
		insert := &db.Fund{
			Code:       fund.Code,
			Name:       fund.Name,
			Status:     db.FundStatusOn,
			Createtime: utils.Now(),
			Updatetime: utils.Now(),
		}
		insertMany = append(insertMany, insert)
	}

	res, err := db.GetFundCol().InsertMany(c, insertMany)
	if err != nil {
		logrus.Errorf("err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "insert fail")
		return
	}

	c.JSON(http.StatusOK, res.InsertedIDs)
	return
}

func GetFundList(c *gin.Context) {
	rsp := make([]db.Fund, 0)
	cur, err := db.GetFundCol().Find(c, bson.D{{}})
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
