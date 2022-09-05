package model

import "github.com/shopspring/decimal"

type TicketEntity struct {
	BaseEntity
	Name        string          `gorm:"column:name;NOT NULL" json:"name"`                                        // 票品名称
	TypeID      int             `gorm:"column:type_id;NOT NULL" json:"typeID"`                                   // 票品类型
	Count       int             `gorm:"column:count;default:1;NOT NULL" json:"count"`                            // 单票库存
	Number      int             `gorm:"column:number;default:0" json:"number"`                                   // 已经销售出去的票数量(仅仅是下单的,超时未支付的会回退)
	PayNumber   int             `gorm:"column:pay_number;default:0" json:"payNumber"`                            // 支付成功的数量(暂时不考虑退票)
	Price       decimal.Decimal `gorm:"type:type:decimal(10,2);column:price;default:0.00;NOT NULL" json:"price"` // 价格
	Status      int             `gorm:"column:status;default:0;NOT NULL" json:"status"`                          // 状态0表上架,1表示下架
	Sort        int             `gorm:"column:sort;default:0;NOT NULL" json:"sort"`                              // 排序
	Description string          `gorm:"column:description" json:"description"`                                   // 描述
	IsLimit     int             `gorm:"column:is_limit;default:0;NOT NULL" json:"isLimit"`                       // 是否限购0表示不限购,1表示限购
	LimitCount  int             `gorm:"column:limit_count" json:"limitCount"`                                    // 限够数量
	ActivityID  int             `gorm:"column:activity_id;NOT NULL" json:"activityID"`                           // 活动ID
	IsRealName  int             `gorm:"column:is_real_name;default:0;NOT NULL" json:"isRealName"`                // 是否实名0表示不实名,1表示实名
}

func (m *TicketEntity) TableName() string {
	return "ticket"
}
