package model

import (
	"database/sql"
	"time"
)

type VjInfoEntity struct {
	BaseEntity
	VjName           string        `gorm:"vj_name" json:"vjName"`                      // vj名称
	ActivityId       int64         `gorm:"activity_id" json:"activityId"`              // 关联活动id
	TicketId         int64         `gorm:"ticket_id" json:"ticketId"`                  // 活动票品id
	SponsorId        int64         `gorm:"sponsor_id" json:"sponsorId"`                // 主办方id
	Type             sql.NullInt64 `gorm:"type" json:"type"`                           // vj类型:0是免费,1是付费
	DueTime          time.Time     `gorm:"due_time" json:"dueTime"`                    // 过期时间
	Title            *string       `gorm:"title" json:"title"`                         // 标题
	GroupName1       *string       `gorm:"group_name1" json:"groupName1"`              // 分组一名称
	GroupName2       *string       `gorm:"group_name2" json:"groupName2"`              // 分组二名称
	ThemeId          int64         `gorm:"theme_id" json:"themeId"`                    // 关联vj_theme表中主键id
	Level            int64         `gorm:"level" json:"level"`                         // 比赛级别：比如32强、16强、8强、4强、2强
	PrimaryBg        *string       `gorm:"primary_bg" json:"primaryBg"`                // 主背景图地址
	Bg32             *string       `gorm:"bg32" json:"bg32"`                           // 32强背景图
	Bg16             *string       `gorm:"bg16" json:"bg16"`                           // 16强背景图
	LogoUrl          *string       `gorm:"logo_url" json:"logoUrl"`                    // logo图地址
	AgainstPlanUrl   *string       `gorm:"against_plan_url" json:"againstPlanUrl"`     // 对阵图图地址
	AgainstFigureUrl *string       `gorm:"against_figure_url" json:"againstFigureUrl"` // 对抗图地址
}

func (t *VjInfoEntity) TableName() string {
	return "vj_info"
}
