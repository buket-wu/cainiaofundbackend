package fund

type AddFundReq struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type GetFundReq struct {
	Code string `json:"code" form:"code" binding:"required"`
}
