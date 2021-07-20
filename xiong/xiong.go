package xiong

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	Prefix           = "https://api.doctorxiong.club/v1"
	UrlGetFund       = "/fund"
	UrlGetFundDetail = "/fund/detail"
)

func GetFund(req *GetFundReq) ([]*Fund, error) {
	rsp := &GetFundRsq{}
	hRsp, err := resty.New().R().SetQueryParams(map[string]string{"code": req.Code}).SetResult(rsp).Get(Prefix + UrlGetFund)
	if err != nil {
		logrus.Errorf("err:%v", err)
		return nil, err
	}

	if hRsp == nil || hRsp.StatusCode() != http.StatusOK {
		return nil, errors.New("get fund fail")
	}

	return rsp.Data, nil
}

func GetFundDetail(req *GetFundDetailReq) (*FundDetail, error) {
	rsp := &GetFundDetailRsq{}
	hRsp, err := resty.New().R().SetQueryParams(map[string]string{"code": req.Code}).SetResult(rsp).Get(Prefix + UrlGetFundDetail)
	if err != nil {
		logrus.Errorf("err:%v", err)
		logrus.Infof("hRsp:%v", hRsp)
		return nil, err
	}

	if hRsp == nil || hRsp.StatusCode() != http.StatusOK {
		return nil, errors.New("get fund fail")
	}

	return rsp.Data, nil
}
