package model

type VjFreeHistoryEntity struct {
	BaseEntity
	VjFreeId int64   `gorm:"vj_free_id" json:"vjFreeId"` // 关联表vj_free的id
	Free     float64 `gorm:"free" json:"free"`           // 本次充值充值费用
	Total    int64   `gorm:"total" json:"total"`         // 本次充值的次数
}

func (t *VjFreeHistoryEntity) TableName() string {
	return "vj_free_history"
}
