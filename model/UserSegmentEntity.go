package model

type UserSegmentEntity struct {
	BaseEntity
	ActivityId      int64  `gorm:"activity_id" json:"activityId"`            // 关联活动表(activity)主键id
	NumberSegmentId int64  `gorm:"number_segment_id" json:"numberSegmentId"` // 关联抽号表(umber_segment)主键id
	Index           int64  `gorm:"index" json:"index"`                       // 抽到号码数字
	UserId          int64  `gorm:"user_id" json:"userId"`                    // 小程序用户id
	NickName        string `gorm:"nick_name" json:"nickName"`                // 小程序用户昵称
	Mobile          string `gorm:"mobile" json:"mobile"`                     // 小程序用户手机号码
}

func (t *UserSegmentEntity) TableName() string {
	return "user_segment"
}
