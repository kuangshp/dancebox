package model

type OrderEntity struct {
	BaseEntity
	OrderSn           string  `gorm:"order_sn" json:"orderSn"`
	TradeNo           string  `gorm:"trade_no" json:"tradeNo"`                      // 支付交易号(微信支付的交易号)
	PayMethod         int     `gorm:"pay_method" json:"payMethod"`                  // 支付方式0是微信,1是支付宝
	UserId            int     `gorm:"user_id" json:"userId"`                        // 交易的用户ID(小程序)
	ActivityId        int     `gorm:"activity_id" json:"activityId"`                // 订单关联活动的ID
	ActivityEndTime   string  `gorm:"activity_end_time" json:"activityEndTime"`     // 活动结束时间
	PaidAmount        float64 `gorm:"paid_amount" json:"paidAmount"`                // 支付金额(总金额)
	StartPaidAt       string  `gorm:"start_paid_at" json:"startPaidAt"`             // 开始支付时间
	EndPaidAt         string  `gorm:"end_paid_at" json:"endPaidAt"`                 // 支付成功时间
	Status            int     `gorm:"status" json:"status"`                         // 支付状态,0表示未支付,1表示已支付,2表示支付失败,3表示取消订单,4表示退款,5表示核销,6表示在支付等待支付回调
	Remark            string  `gorm:"remark" json:"remark"`                         // 备注信息
	CancelOrderReason string  `gorm:"cancel_order_reason" json:"cancelOrderReason"` // 取消订单原因
	SponsorId         int     `gorm:"sponsor_id" json:"sponsorId"`                  // 订单关联主办方ID
	ActivityTypeId    int     `gorm:"activity_type_id" json:"activityTypeId"`       // 活动类型id
	TicketTotal       int     `gorm:"ticket_total" json:"ticketTotal"`              // 该订单下票数量
	DeletedAt         string  `gorm:"deleted_at" json:"deletedAt"`                  // 软删除时间
	ActivityStartTime string  `gorm:"activity_start_time" json:"activityStartTime"` // 活动开始时间
	RefundStatus      int     `gorm:"refund_status" json:"refundStatus"`            // 退款状态
	RefundTicketTotal int     `gorm:"refund_ticket_total" json:"refundTicketTotal"` // 退票数量
}

func (*OrderEntity) TableName() string {
	return "order"
}
