package model

type SponsorPlugEntity struct {
	BaseEntity
	Title string `gorm:"title" json:"title"` // 插件名称
	Type  int64  `gorm:"type" json:"type"`   // 分组
}

func (t *SponsorPlugEntity) TableName() string {
	return "sponsor_plug"
}
