package model

type SponsorAccountRoleEntity struct {
	BaseEntity
	AccountId int64 `gorm:"account_id" json:"accountId"` // 关联主办方账号表的ID
	SponsorId int64 `gorm:"sponsor_id" json:"sponsorId"` // 关联主办方表的ID
	RoleId    int64 `gorm:"role_id" json:"roleId"`       // 关联角色表id
}

func (t *SponsorAccountRoleEntity) TableName() string {
	return "sponsor_account_role"
}
