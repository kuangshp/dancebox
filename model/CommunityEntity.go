package model

import (
	"database/sql"
	"time"
)

type CommunityEntity struct {
	BaseEntity
	Name          string        `gorm:"name" json:"name"`                    // 名称
	Describe      string        `gorm:"describe" json:"describe"`            // 描述
	CommunityType int64         `gorm:"community_type" json:"communityType"` // 分类,1表示兴趣群2表示区域群
	ProvinceId    sql.NullInt64 `gorm:"province_id" json:"provinceId"`       // 省份Id
	CityId        sql.NullInt64 `gorm:"city_id" json:"cityId"`               // 城市Id
	QrGroupUrl    string        `gorm:"qr_group_url" json:"qrGroupUrl"`      // 群二维码
	ValidTime     time.Time     `gorm:"valid_time" json:"validTime"`         // 有效期
	WechatUrl     string        `gorm:"wechat_url" json:"wechatUrl"`         // 微信二维码
	IsHidden      int64         `gorm:"is_hidden" json:"isHidden"`           // 状态,1表示显示2表示隐藏
	IsFull        int64         `gorm:"is_full" json:"isFull"`               // 是否满,1表示没有,2表示满群
	MaxCount      int64         `gorm:"max_count" json:"maxCount"`           // 默认满员数量
	CurrentCount  int64         `gorm:"current_count" json:"currentCount"`   // 当前加群人数
	Sort          int64         `gorm:"sort" json:"sort"`                    // 排序
}

func (t *CommunityEntity) TableName() string {
	return "community"
}
