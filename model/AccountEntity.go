package model

import (
	"time"
)

type AccountEntity struct {
	BaseEntity
	Username        string              `json:"username" gorm:"column:username;NOT NULL"`                                       // 用户名
	Password        string              `json:"password" gorm:"column:password;NOT NULL"`                                       // 密码
	Name            string              `json:"name" gorm:"column:name"`                                                        // 用户真实姓名
	Status          int                 `json:"status" gorm:"column:status;default:1"`                                          // 状态
	Platform        int                 `json:"platform" gorm:"column:platform;default:1"`                                      // 平台:1为主办方,2为运营管理
	IsSuper         int                 `json:"isSuper" gorm:"column:is_super;default:0;NOT NULL"`                              // 是否为超级管理员1表示是,0表示不是
	SponsorID       int                 `json:"sponsorId" gorm:"column:sponsor_id"`                                             // 是否成功发送邮件0表示没有，1表示发送成功
	LastLoginIP     string              `json:"lastLoginIp" gorm:"column:last_login_ip"`                                        // 最后登录id
	LastLoginTime   time.Time           `json:"lastLoginTime" gorm:"column:last_login_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 最后登录时间
	OpenID          string              `json:"openId" gorm:"column:open_id"`                                                   // 小程序助手的openId
	Unionid         string              `json:"unionid" gorm:"column:unionid"`                                                  // 微信唯一吗
	SessionKey      string              `json:"sessionKey" gorm:"column:session_key"`                                           // 微信的sessionKey
	AccountRoleList []AccountRoleEntity `json:"accountRole"`
	RoleList        []*RoleEntity       `json:"roleList" gorm:"many2many:account_role;foreignKey:ID;joinForeignKey:accountId;joinReferences:roleId"`
}

func (m *AccountEntity) TableName() string {
	return "account"
}
