package model

type RoleAccessEntity struct {
	BaseEntity
	RoleID   int64  `json:"roleId" gorm:"column:role_id"`      // 角色id
	AccessID int64  `json:"accessId" gorm:"column:access_id"`  // 资源id
	Type     string `json:"type" gorm:"column:type;default:1"` // 类型,type=1表示菜单,type=2表示接口
}

func (m *RoleAccessEntity) TableName() string {
	return "role_access"
}
