package model

type AccessEntity struct {
	BaseEntity
	ModuleName  string `json:"moduleName" gorm:"column:module_name"`                 // 模块名称
	Type        int    `json:"type" gorm:"column:type"`                              // 类型
	ActionName  string `json:"actionName" gorm:"column:action_name"`                 // 操作名称
	Icon        string `json:"icon" gorm:"column:icon"`                              // 小图标
	Url         string `json:"url" gorm:"column:url"`                                // url地址
	Method      string `json:"method" gorm:"column:method"`                          // 请求方式
	ModuleID    int    `json:"moduleId" gorm:"column:module_id;default:-1;NOT NULL"` // 父模块id
	Sort        int    `json:"sort" gorm:"column:sort;default:1;NOT NULL"`           // 排序
	Description string `json:"description" gorm:"column:description"`                // 描述
	Status      int    `json:"status" gorm:"column:status;default:1"`                // 状态,1表示正常,0表示禁用
}

func (m *AccessEntity) TableName() string {
	return "access"
}
