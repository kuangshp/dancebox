package model

import "time"

type WeiXinUserTicketEntity struct {
	BaseEntity
	UserId           int64     `gorm:"user_id" json:"userId"`                        // 小程序用户id
	WeiXinRealUserId int64     `gorm:"wei_xin_real_user_id" json:"weiXinRealUserId"` // 关联wei_xin_real_user表的id
	ActivityId       int64     `gorm:"activity_id" json:"activityId"`                // 活动id
	ActivityTitle    string    `gorm:"activity_title" json:"activityTitle"`          // 活动标题
	ActivityEndTime  time.Time `gorm:"activity_end_time" json:"activityEndTime"`     // 活动结束时间
	TicketName       string    `gorm:"ticket_name" json:"ticketName"`                // 票品名称
	TicketType       string    `gorm:"ticket_type" json:"ticketType"`                // 票品类型
	OrderListId      int64     `gorm:"order_list_id" json:"orderListId"`             // 在订单列表中对应的id
	TicketId         int64     `gorm:"ticket_id" json:"ticketId"`                    // 票品id
}

func (t *WeiXinUserTicketEntity) TableName() string {
	return "wei_xin_user_ticket"
}
