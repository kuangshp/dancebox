package model

type SponsorAccessEntity struct {
	BaseEntity
	AccessId        int64  `gorm:"access_id" json:"accessId"`                // 关联资源表表的ID
	Mobile          string `gorm:"mobile" json:"mobile"`                     // 手机号码
	ParentSponsorId int64  `gorm:"parent_sponsor_id" json:"parentSponsorId"` // 关联主号主办方id
}

func (t *SponsorAccessEntity) TableName() string {
	return "sponsor_access"
}
