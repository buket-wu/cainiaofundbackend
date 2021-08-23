package fund

type AddFundReq struct {
	Codes string `json:"codes" form:"codes" binding:"required"`
}

type AddFundResp struct {
	SuccessCodeList []string `json:"successCodeList"`
}
