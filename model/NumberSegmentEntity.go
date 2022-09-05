package model

import "time"

type NumberSegmentEntity struct {
	BaseEntity
	Name       string    `gorm:"name" json:"name"`              // 抽号名称
	PrimaryBg  string    `gorim:"primary_bg" json:"primaryBg"`  // 背景图
	ActivityId int64     `gorm:"activity_id" json:"activityId"` // 关联活动表(activity)主键id
	StartTime  time.Time `gorm:"start_time" json:"startTime"`   // 开始抽号时间
	EndTime    time.Time `gorm:"end_time" json:"endTime"`       // 结束抽号时间
	Rule       int64     `gorm:"rule" json:"rule"`              // 抽号规则1表示随机，2表示顺序
	StartNum   int64     `gorm:"start_num" json:"startNum"`     // 开始抽号数字
	EndNum     int64     `gorm:"end_num" json:"endNum"`         // 结束抽号数字
	TicketId   int64     `gorm:"ticket_id" json:"ticketId"`     // 抽号群体关联到ticket表中的id
	TicketName string    `gorm:"ticket_name" json:"ticketName"` // 抽号群体名称
	TypeId     int64     `gorm:"type_id" json:"typeId"`         // 票类型id
	TypeName   string    `gorm:"type_name" json:"typeName"`     // 票类型名称
}

func (t *NumberSegmentEntity) TableName() string {
	return "number_segment"
}
