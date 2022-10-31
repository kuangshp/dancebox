package model

import "time"

// QrGroupUrlEntity 具体图片地址的
type QrGroupUrlEntity struct {
	BaseEntity
	QrGroupId int64     `gorm:"qr_group_id" json:"qrGroupId"` // 管理二维码群的外键id
	QrUrl     string    `gorm:"qr_url" json:"qrUrl"`          // 加群的二维码
	DueTime   time.Time `gorm:"due_time" json:"dueTime"`      // 二维码过期时间
	MaxNum    int64     `gorm:"max_num" json:"maxNum"`        // 群最大人数
	CurNum    int64     `gorm:"cur_num" json:"curNum"`        // 当前人数
}

func (t *QrGroupUrlEntity) TableName() string {
	return "qr_group_url"
}
