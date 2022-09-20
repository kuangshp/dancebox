package model

type DanceEntity struct {
	BaseEntity
	Name        string `json:"name" gorm:"column:name;comment:舞蹈名称"`                         // 舞蹈名称
	Sort        int    `json:"sort" gorm:"column:sort;default:1;NOT NULL;comment:排序"`        // 排序
	Description string `json:"description" gorm:"column:description;comment:描述"`             // 描述
	Status      int    `json:"status" gorm:"column:status;default:1;comment:状态,1表示正常,0表示禁用"` // 状态,1表示正常,0表示禁用
}

func (m *DanceEntity) TableName() string {
	return "dance"
}
