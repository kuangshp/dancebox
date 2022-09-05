package model

import "database/sql"

type VjPromotionEntity struct {
	BaseEntity
	Level     int64         `gorm:"level" json:"level"`          // 比赛级别：比如32强、16强、8强、4强、2强
	Name      *string       `gorm:"name" json:"name"`            // 参赛选手姓名
	Url       *string       `gorm:"url" json:"url"`              // 参赛选手头像地址
	ParentId  int64         `gorm:"parent_id" json:"parentId"`   // 晋级父节点id
	Index     int64         `gorm:"index" json:"index"`          // 参赛选手的位置
	InfoId    int64         `gorm:"info_id" json:"infoId"`       // 关联vj_info表主键id
	IsUpgrade sql.NullInt64 `gorm:"is_upgrade" json:"IsUpgrade"` // 是否晋升0表示没有,1表示晋升了
}

func (t *VjPromotionEntity) TableName() string {
	return "vj_promotion"
}
