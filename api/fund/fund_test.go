package fund

import (
	"cainiaofundbackend/db"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"testing"
)

func TestFund_GetCode(t *testing.T) {
	ctx := context.Background()
	rsp := make([]db.Fund, 0)
	cur, err := db.FundCol.Find(ctx, bson.D{{}})
	if err != nil {
		logrus.Errorf("err:%v", err)
		return
	}

	err = cur.All(ctx, &rsp)
	if err != nil {
		logrus.Errorf("err:%v", err)
	}

	arr := make([]string, 0)
	for _, fund := range rsp {
		arr = append(arr, fund.Code)
	}

	// 004075,008507,008314,070032,001410,006257,270002,180020,320007,005962,002190,000939,002251
	logrus.Infof("code:str:%v", strings.Join(arr, ","))
}
