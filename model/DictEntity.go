package model

type DictEntity struct {
	BaseEntity
	Category       string `json:"category" gorm:"column:category;NOT NULL"`      // 分类名称
	Description    string `json:"description" gorm:"column:description"`         // 描述
	Label          string `json:"label" gorm:"column:label;NOT NULL"`            // label值
	TicketTypeRate string `json:"ticketTypeRate" gorm:"column:ticket_type_rate"` // 票种费率
	Value          string `json:"value" gorm:"column:value"`                     // value值
}

func (m *DictEntity) TableName() string {
	return "dict"
}
