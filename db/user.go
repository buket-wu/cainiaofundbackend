package db

type User struct {
	Openid     string `bson:"openid" json:"openid"`
	Username   string `bson:"username" json:"username"`     // 用户名
	Createtime uint32 `bson:"createtime" json:"createtime"` // 创建时间
	Updatetime uint32 `bson:"updatetime" json:"updatetime"` // 更新时间
}
