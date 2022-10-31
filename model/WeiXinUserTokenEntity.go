package model

import "time"

type WeiXinUserTokenEntity struct {
	BaseEntity
	UserId     int64     `gorm:"user_id" json:"userId"`         // 关联核销账号的ID
	Token      string    `gorm:"token" json:"token"`            // token
	ExpireTime time.Time `gorm:"expire_time" json:"expireTime"` // 失效时间
}

func (t *WeiXinUserTokenEntity) TableName() string {
	return "wei_xin_user_token"
}
