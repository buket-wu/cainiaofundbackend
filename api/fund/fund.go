package fund

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/db/dbtools"
	"cainiaofundbackend/extend/utils"
	xiong2 "cainiaofundbackend/extend/xiong"
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

	fundList := make([]db.Fund, 0)
	err := dbtools.GetMany(c, db.GetFundCol(), &fundList, bson.M{"code": bson.M{"$in": codeArr}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "get local fund fail")
		return
	}

	newCodeList := make([]string, 0)
	for _, fund := range fundList {
		if _, ok := codeMap[fund.Code]; !ok {
			newCodeList = append(newCodeList)
		}
	}

	if len(newCodeList) == 0 {
		c.JSON(http.StatusOK, resp)
		return
	}

	xReq := &xiong2.GetFundReq{
		Code: strings.Join(newCodeList, ","),
	}

	xiongFundList, err := xiong2.GetFund(xReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "get fund fail")
		return
	}

	insertMany := make([]interface{}, 0)
	for _, fund := range xiongFundList {
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
	err := dbtools.GetMany(c, db.GetFundCol(), &rsp, bson.M{})
	if err != nil {
		logrus.Errorf("err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "query fail")
		return
	}

	c.JSON(http.StatusOK, rsp)
	return
}
