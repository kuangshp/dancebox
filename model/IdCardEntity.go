package model

type IdCardEntity struct {
	BaseEntity
	CardNo   string `gorm:"card_no" json:"cardNo"`    // 身份证号码
	Name     string `gorm:"name" json:"name"`         // 姓名
	Address  string `gorm:"address" json:"address"`   // 地址
	Birthday string `gorm:"birthday" json:"birthday"` // 出生年月
	Sex      string `gorm:"sex" json:"sex"`           // 性别
}

func (t *IdCardEntity) TableName() string {
	return "id_card"
}
