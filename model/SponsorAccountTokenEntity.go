package model

import "time"

type SponsorAccountTokenEntity struct {
	BaseEntity
	AccountId  int64     `gorm:"account_id" json:"accountId"`   // 关联主办方账号表的ID
	SponsorId  int64     `gorm:"sponsor_id" json:"sponsorId"`   // 关联主办方表的ID
	Token      string    `gorm:"token" json:"token"`            // token
	ExpireTime time.Time `gorm:"expire_time" json:"expireTime"` // 失效时间
}

func (t *SponsorAccountTokenEntity) TableName() string {
	return "sponsor_account_token"
}
