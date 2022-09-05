package model

type IdCardHistoryEntity struct {
	BaseEntity
	SponsorId int    `gorm:"sponsor_id" json:"sponsorId"` // 主办方ID
	Status    int    `gorm:"status" json:"status"`        // 调用状态
	CardNo    string `gorm:"card_no" json:"cardNo"`       // 身份证号码
	Name      string `gorm:"name" json:"name"`            // 姓名
	Address   string `gorm:"address" json:"address"`      // 地址
	Birthday  string `gorm:"birthday" json:"birthday"`    // 出生年月
	Sex       string `gorm:"sex" json:"sex"`              // 性别
	Type      int    `gorm:"type" json:"type"`            // 类型,0表示本地数据库读取，1表示远程接口获取
}

func (*IdCardHistoryEntity) TableName() string {
	return "id_card_history"
}
