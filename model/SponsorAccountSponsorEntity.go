package model

type SponsorAccountSponsorEntity struct {
	BaseEntity
	Mobile                 string `gorm:"mobile" json:"mobile"`                                    // 手机号码
	ParentSponsorId        int64  `gorm:"parent_sponsor_id" json:"parentSponsorId"`                // 关联主号主办方id
	ParentSponsorAccountId int64  `gorm:"parent_sponsor_account_id" json:"parentSponsorAccountId"` // 关联主办方账号表id
	NickName               string `gorm:"nick_name" json:"nickName"`                               // 昵称
}

func (t *SponsorAccountSponsorEntity) TableName() string {
	return "sponsor_account_sponsor"
}
