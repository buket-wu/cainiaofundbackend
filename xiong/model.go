package xiong

type GetFundDetailReq struct {
	Code string `json:"code"`
}

type GetFundDetailRsq struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    *FundDetail `json:"data"`
}

type FundDetail struct {
	Code                  string     `json:"code"` // 基金代码
	Name                  string     `json:"name"` // 基金名称
	Type                  string     `json:"type"` // 基金类型
	NetWorth              float32    `json:"netWorth"`
	ExpectWorth           float32    `json:"expectWorth"`
	TotalWorth            float32    `json:"totalWorth"`
	ExpectGrowth          string     `json:"expectGrowth"`
	DayGrowth             string     `json:"dayGrowth"`
	LastWeekGrowth        string     `json:"lastWeekGrowth"`
	LastMonthGrowth       string     `json:"lastMonthGrowth"`
	LastThreeMonthsGrowth string     `json:"lastThreeMonthsGrowth"`
	LastSixMonthsGrowth   string     `json:"lastSixMonthsGrowth"`
	LastYearGrowth        string     `json:"lastYearGrowth"`
	BuyMin                string     `json:"buyMin"`
	BuySourceRate         string     `json:"buySourceRate"`
	BuyRate               string     `json:"buyRate"`
	Manager               string     `json:"manager"`
	FundScale             string     `json:"fundScale"`
	NetWorthDate          string     `json:"netWorthDate"`
	ExpectWorthDate       string     `json:"expectWorthDate"`
	NetWorthData          [][]string `json:"netWorthData"`
}

type GetFundReq struct {
	Code string `json:"code"`
}

type GetFundRsq struct {
	Code    uint    `json:"code"`
	Message string  `json:"message"`
	Data    []*Fund `json:"data"`
}

type Fund struct {
	Code                  string  `json:"code"` // 基金代码
	Name                  string  `json:"name"` // 基金名称
	Type                  string  `json:"type"` // 基金类型
	NetWorth              float32 `json:"netWorth"`
	ExpectWorth           float32 `json:"expectWorth"`
	ExpectGrowth          string  `json:"expectGrowth"`
	DayGrowth             string  `json:"dayGrowth"`
	LastWeekGrowth        string  `json:"lastWeekGrowth"`
	LastMonthGrowth       string  `json:"lastMonthGrowth"`
	LastThreeMonthsGrowth string  `json:"lastThreeMonthsGrowth"`
	LastSixMonthsGrowth   string  `json:"lastSixMonthsGrowth"`
	LastYearGrowth        string  `json:"lastYearGrowth"`
	NetWorthDate          string  `json:"netWorthDate"`
	ExpectWorthDate       string  `json:"expectWorthDate"`
}
