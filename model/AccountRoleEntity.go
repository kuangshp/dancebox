package model

type AccountRoleEntity struct {
	BaseEntity
	AccountEntityId int `json:"accountId" gorm:"column:account_id;NOT NULL"` // 账号id
	RoleEntityId    int `json:"roleId" gorm:"column:role_id;NOT NULL"`       // 角色id
}

func (m *AccountRoleEntity) TableName() string {
	return "account_role"
}
