package model

type RoleEntity struct {
	BaseEntity
	Title       string              `json:"title" gorm:"column:title;NOT NULL"`           // 角色名称
	Description string              `json:"description" gorm:"column:description"`        // 角色描述
	Status      int                 `json:"status" gorm:"column:status;default:1"`        // 状态1表示正常,0表示不正常
	IsDefault   int                 `json:"isDefault" gorm:"column:is_default;default:0"` // 针对主办方自己创建,默认开通的角色,1表示开通,0表示不开通
	Accounts    []AccountRoleEntity `json:"accounts"`
	// many2many:user_profiles 定义中间表名:many2many:user_profiles
	// foreignKey:ID 使用id作为外键
	// joinForeignKey:profileId" 自定义中间表的外键名称
	// joinReferences 这个要加上反向引用名称
	AccountList []*AccountEntity `json:"accountsList" gorm:"many2many:account_role;foreignKey:ID;joinForeignKey:roleId;joinReferences:accountId"`
}

func (m *RoleEntity) TableName() string {
	return "role"
}
