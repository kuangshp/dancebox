package model

type QrGroupEntity struct {
	BaseEntity
	IsDisabled int64  `gorm:"is_disabled" json:"isDisabled"` // 是否禁用,1表示禁用,0表示正常
	ActivityId int64  `gorm:"activity_id" json:"activityId"` // 关联的活动id
	TicketId   int64  `gorm:"ticket_id" json:"ticketId"`     // 关联的票id
	WxCard     string `gorm:"wx_card" json:"wxCard"`         // 微信二维码名片
	AccountId  int64  `gorm:"account_id" json:"accountId"`   // 关联当前登录账号id
	SponsorId  int64  `gorm:"sponsor_id" json:"sponsorId"`   // 主办方id
}

func (t *QrGroupEntity) TableName() string {
	return "qr_group"
}
