package model

type VjFreeEntity struct {
	BaseEntity
	SponsorId int64   `gorm:"sponsor_id" json:"sponsorId"` // 主办方id
	Free      float64 `gorm:"free" json:"free"`            // 总充值费用
	AllTotal  int64   `gorm:"all_total" json:"allTotal"`   // 总次数
	UseTotal  int64   `gorm:"use_total" json:"useTotal"`   // 已经使用
}

func (t *VjFreeEntity) TableName() string {
	return "vj_free"
}
