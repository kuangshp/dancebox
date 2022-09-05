package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type SponsorAccountEntity struct {
	BaseEntity
	Username      string    `gorm:"username" json:"username"`             // 用户名
	Password      string    `gorm:"password" json:"password"`             // 密码
	NickName      string    `gorm:"nick_name" json:"nickName"`            // 微信昵称
	AvatarUrl     string    `gorm:"avatar_url" json:"avatarUrl"`          // 微信头像
	Mobile        string    `gorm:"mobile" json:"mobile"`                 // 手机号码
	Status        int64     `gorm:"status" json:"status"`                 // 状态(1正常,0禁用)
	SponsorId     int64     `gorm:"sponsor_id" json:"sponsorId"`          // 主办方Id(主要的)
	LastLoginIp   string    `gorm:"last_login_ip" json:"lastLoginIp"`     // 最后登录id
	LastLoginTime time.Time `gorm:"last_login_time" json:"lastLoginTime"` // 最后登录时间
	OpenId        string    `gorm:"open_id" json:"openId"`                // 小程序助手的openId
	Unionid       string    `gorm:"unionid" json:"unionid"`               // 微信唯一吗
	SessionKey    string    `gorm:"session_key" json:"sessionKey"`        // 微信的sessionKey
}

func (t *SponsorAccountEntity) TableName() string {
	return "sponsor_account"
}
