package job

import (
	"cainiaofundbackend/db"
	"cainiaofundbackend/db/dbtools"
	"cainiaofundbackend/extend/message"
	"cainiaofundbackend/extend/routine"
	utils2 "cainiaofundbackend/extend/utils"
	xiong2 "cainiaofundbackend/extend/xiong"
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
	err := dbtools.GetMany(ctx, db.GetFundCol(), &fundList, bson.M{"status": db.FundStatusOn})
	codeArr := make([]string, len(fundList))
	for _, fund := range fundList {
		codeArr = append(codeArr, fund.Code)
	}

	var whereTime int64
	if nowWeekday == time.Monday {
		whereTime = utils2.GetFirstDateOfWeek(now).Unix()
	} else {
		whereTime = utils2.GetLastWeekFirstDate(now).Unix()
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})
	fundTrendList := make([]db.FundTrend, 0)
	err = dbtools.GetMany(ctx, db.GetFundTrendCol(), &fundTrendList, bson.M{
		"isMonday":   1,
		"isDayLast":  1,
		"code":       bson.M{"$in": codeArr},
		"createtime": bson.M{"$gt": whereTime}}, findOptions)
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

	req := xiong2.GetFundReq{
		Code: strings.Join(codeArr, ","),
	}
	rsp, err := xiong2.GetFund(&req)
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
			SpecGrowth = ((fund.NetWorth - lastTrend.NetWorth) / lastTrend.NetWorth) * 100
			if SpecGrowth >= db.RemindMinGrowth {
				record := &db.RemindRecord{
					Code:        fund.Code,
					UserOpenid:  "60ebabcc2a40500ff3040966",
					NetWorth:    fund.NetWorth,
					ExpectWorth: fund.ExpectWorth,
					SpecGrowth:  SpecGrowth,
					Createtime:  utils2.Now(),
					Updatetime:  utils2.Now(),
				}
				insertRecord = append(insertRecord, record)

				// 发送邮件提醒
				err = routine.Pool.Submit(func() {
					err = message.NewMessageDrive(message.DriveTypeWechatDefault).Send(record.UserOpenid, "xxx")
					if err != nil {
						logrus.Errorf("send messge fail; err:%v", err)
					}
				})
				if err != nil {
					logrus.Errorf("send messge fail; err:%v", err)
				}
			}
		} else {
			SpecGrowth = 0
		}

		insert := &db.FundTrend{
			Code:         fund.Code,
			Name:         fund.Name,
			NetWorth:     fund.NetWorth,
			ExpectWorth:  fund.ExpectWorth,
			IsMonday:     utils2.Bool2Uint32(nowWeekday == time.Monday),
			IsDayLast:    utils2.Bool2Uint32(nowHour >= 15),
			DayGrowth:    fund.DayGrowth,
			ExpectGrowth: fund.ExpectGrowth,
			SpecGrowth:   SpecGrowth,
			Createtime:   utils2.Now(),
			Updatetime:   utils2.Now(),
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
