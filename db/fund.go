package db

const (
	FundStatusOn = 1
)

const RemindMinGrowth = 5

type Fund struct {
	Code       string `bson:"code" json:"code"`             // 基金代码
	Name       string `bson:"name" json:"name"`             // 基金名称
	Status     uint32 `bson:"status" json:"status"`         // 启用状态
	Createtime uint32 `bson:"createtime" json:"createtime"` // 创建时间
	Updatetime uint32 `bson:"updatetime" json:"updatetime"` // 更新时间
}

type FundTrend struct {
	Code         string  `bson:"code" json:"code"`                 // 基金代码
	Name         string  `bson:"name" json:"name"`                 // 基金名称
	NetWorth     float32 `bson:"netWorth" json:"netWorth"`         // 净值
	ExpectWorth  float32 `bson:"expectWorth" json:"expectWorth"`   // 估算净值
	IsMonday     uint32  `bson:"isMonday" json:"isMonday"`         // 是否是周一
	IsDayLast    uint32  `bson:"isDayLast" json:"isDayLast"`       // 是否为日计算
	DayGrowth    string  `bson:"dayGrowth" json:"dayGrowth"`       // 日涨幅
	ExpectGrowth string  `bson:"expectGrowth" json:"expectGrowth"` // 估算净值
	SpecGrowth   float32 `bson:"specGrowth" json:"specGrowth"`     // 对比周一涨幅（周一则对比上周一）
	Createtime   uint32  `bson:"createtime" json:"createtime"`     // 创建时间
	Updatetime   uint32  `bson:"updatetime" json:"updatetime"`     // 更新时间
}

type FundBelong struct {
	Code       string  `bson:"code" json:"code"`             // 基金代码
	UserID     string  `bson:"userID" json:"userID"`         // 用户id
	SpecGrowth float32 `bson:"specGrowth" json:"specGrowth"` // 发送提醒百分比
	Createtime uint32  `bson:"createtime" json:"createtime"` // 创建时间
	Updatetime uint32  `bson:"updatetime" json:"updatetime"` // 更新时间
}

type RemindRecord struct {
	Code         string  `bson:"code" json:"code"`                 // 基金代码
	UserOpenid   string  `bson:"userOpenid" json:"userOpenid"`     // 用户id
	NetWorth     float32 `bson:"netWorth" json:"netWorth"`         // 净值（上一个工作日l）
	ExpectWorth  float32 `bson:"expectWorth" json:"expectWorth"`   // 估算净值（当天）
	ExpectGrowth string  `bson:"expectGrowth" json:"expectGrowth"` // 估算涨幅（当天）
	SpecGrowth   float32 `bson:"specGrowth" json:"specGrowth"`     // 对比周一涨幅（周一则对比上周一）
	Createtime   uint32  `bson:"createtime" json:"createtime"`     // 创建时间
	Updatetime   uint32  `bson:"updatetime" json:"updatetime"`     // 更新时间
}
