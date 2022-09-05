package model

import "database/sql"

type VjThemeEntity struct {
	BaseEntity
	Name            string        `gorm:"name" json:"name"`                         // 主题名称
	Url             string        `gorm:"url" json:"url"`                           // 主题背景图地址
	IsSystemDefault sql.NullInt64 `gorm:"is_system_default" json:"isSystemDefault"` // 是否系统预设的主题0表示不是，1表示是
}

func (t *VjThemeEntity) TableName() string {
	return "vj_theme"
}
