package xiong

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetFund(t *testing.T) {

}

func TestGetFundDetail(t *testing.T) {
	req := &GetFundDetailReq{
		Code: "004075",
	}

	res, err := GetFundDetail(req)

	jsonStr, _ := json.Marshal(res)
	logrus.Info(string(jsonStr))
	logrus.Info(err)
}
