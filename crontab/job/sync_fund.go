package job

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/utils"
	"cainiaofundbackend/xiong"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type SyncFund struct{}

func (j SyncFund) Run() {

	ctx := getCtx()
	now := time.Now()

	nowHour := now.Hour()
	if nowHour < 9 || nowHour > 16 {
		return
	}
	nowWeekday := now.Weekday()
	if nowWeekday == time.Sunday || nowWeekday == time.Saturday {
		return
	}

	fundList := make([]db.Fund, 0)
	cur, err := db.GetFundCol().Find(ctx, bson.M{"status": db.FundStatusOn})
	if err != nil {
		logrus.Errorf("get fund fail; err:%v", err)
		return
	}
	defer func() {
		err := cur.Close(ctx)
		if err != nil {
			logrus.Errorf("mongo cur fail; err:%v", err)
		}
	}()

	err = cur.All(ctx, &fundList)
	if err != nil {
		logrus.Errorf("cur fund fail; err:%v", err)
		return
	}

	codeArr := make([]string, len(fundList))
	for _, fund := range fundList {
		codeArr = append(codeArr, fund.Code)
	}

	var whereTime int64
	if nowWeekday == time.Monday {
		whereTime = utils.GetFirstDateOfWeek(now).Unix()
	} else {
		whereTime = utils.GetLastWeekFirstDate(now).Unix()
	}

	fundTrendList := make([]db.FundTrend, 0)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})
	cur, err = db.GetFundTrendCol().Find(ctx, bson.M{
		"isMonday":   1,
		"isDayLast":  1,
		"code":       bson.M{"$in": codeArr},
		"createtime": bson.M{"$gt": whereTime}}, findOptions)
	if err != nil {
		logrus.Errorf("get fund trend fail; err:%v", err)
		return
	}
	err = cur.All(ctx, &fundTrendList)
	if err != nil {
		logrus.Errorf("cur fund trend fail; err:%v", err)
		return
	}
	fundTrendMap := make(map[string]db.FundTrend)
	if len(fundTrendList) > 0 {
		for _, fundTrend := range fundTrendList {
			if _, ok := fundTrendMap[fundTrend.Code]; !ok {
				fundTrendMap[fundTrend.Code] = fundTrend
			}
		}
	}

	req := xiong.GetFundReq{
		Code: strings.Join(codeArr, ","),
	}
	rsp, err := xiong.GetFund(&req)
	if err != nil {
		logrus.Errorf("get xiong fund; err:%v", err)
		return
	}

	insertMany := make([]interface{}, 0)
	insertRecord := make([]interface{}, 0)
	for _, fund := range rsp {
		var SpecGrowth float32
		lastTrend, ok := fundTrendMap[fund.Code]
		if ok {
			// todo::判断是否提醒
			SpecGrowth = ((fund.NetWorth - lastTrend.NetWorth) / lastTrend.NetWorth) * 100
			if SpecGrowth <= 5 {
				record := &db.RemindRecord{
					Code:        fund.Code,
					UserOpenid:  "60ebabcc2a40500ff3040966",
					NetWorth:    fund.NetWorth,
					ExpectWorth: fund.ExpectWorth,
					SpecGrowth:  SpecGrowth,
					Createtime:  utils.Now(),
					Updatetime:  utils.Now(),
				}
				insertRecord = append(insertRecord, record)
			}
		} else {
			SpecGrowth = 0
		}

		insert := &db.FundTrend{
			Code:         fund.Code,
			Name:         fund.Name,
			NetWorth:     fund.NetWorth,
			ExpectWorth:  fund.ExpectWorth,
			IsMonday:     utils.Bool2Uint32(nowWeekday == time.Monday),
			IsDayLast:    utils.Bool2Uint32(nowHour >= 15),
			DayGrowth:    fund.DayGrowth,
			ExpectGrowth: fund.ExpectGrowth,
			SpecGrowth:   SpecGrowth,
			Createtime:   utils.Now(),
			Updatetime:   utils.Now(),
		}
		insertMany = append(insertMany, insert)
	}

	_, err = db.GetFundTrendCol().InsertMany(ctx, insertMany)
	if err != nil {
		logrus.Errorf("err:%v", err)
		return
	}

	if len(insertRecord) > 0 {
		_, err = db.GetRemindRecordCol().InsertMany(ctx, insertRecord)
		if err != nil {
			logrus.Errorf("err:%v", err)
			return
		}
	}

	return
}
