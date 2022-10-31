package model

import "time"

type WeiXinRealUserEntity struct {
	BaseEntity
	Name       string    `gorm:"name" json:"name"`               // 参赛名称或真实姓名
	Mobile     string    `gorm:"mobile" json:"mobile"`           // 手机号码
	Birthday   time.Time `gorm:"birthday" json:"birthday"`       // 生日
	UserId     int64     `gorm:"user_id" json:"userId"`          // 关联小程序用户id
	Gender     string    `gorm:"gender" json:"gender"`           // 性别
	Province   string    `gorm:"province" json:"province"`       // 省份
	City       string    `gorm:"city" json:"city"`               // 城市
	TeamName   string    `gorm:"team_name" json:"teamName"`      // 战队名称
	CardNo     string    `gorm:"card_no" json:"cardNo"`          // 身份证号码
	IdCardName string    `gorm:"id_card_name" json:"idCardName"` // 身份证上的姓名
}

func (t *WeiXinRealUserEntity) TableName() string {
	return "wei_xin_real_user"
}
