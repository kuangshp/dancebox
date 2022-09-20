package model

type WeChatOpenUserEntity struct {
	BaseEntity
	OpenId  string `json:"openid" gorm:"column:openid"`
	Unionid string `json:"unionid" gorm:"column:unionid"`
}

func (m *WeChatOpenUserEntity) TableName() string {
	return "wechat_open_user"
}
