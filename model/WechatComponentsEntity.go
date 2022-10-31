package model

import (
	"database/sql"
)

type WechatComponentsEntity struct {
	BaseEntity
	Name     string        `gorm:"name" json:"name"`         // 名称
	Describe string        `gorm:"describe" json:"describe"` // 描述
	Position int64         `gorm:"position" json:"position"` // 组件位置,前端自定义
	Style    int64         `gorm:"style" json:"style"`       // 组件样式1表示一行一，2表示二行三，3表示一行三
	Status   sql.NullInt64 `gorm:"status" json:"status"`     // 状态,0表示隐藏,1表示显示
	Sort     int64         `gorm:"sort" json:"sort"`         // 排序
}

func (t *WechatComponentsEntity) TableName() string {
	return "wechat_components"
}
