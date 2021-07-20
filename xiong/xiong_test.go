package xiong

import (
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
	logrus.Info(res)
	logrus.Info(err)
}
