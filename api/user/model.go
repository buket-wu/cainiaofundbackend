package user

type AddUserReq struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type DelUserReq struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type AddMyFundReq struct {
	Code string `json:"code" form:"code" binding:"required"`
}
