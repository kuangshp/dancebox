package model

type OrderListEntity struct {
	BaseEntity
	TicketNo              string  `gorm:"ticket_no" json:"ticketNo"`                             // 每一张票对应的唯一ID
	TicketId              int     `gorm:"ticket_id" json:"ticketId"`                             // 票品id
	WeiXinRealUserId      int     `gorm:"wei_xin_real_user_id" json:"weiXinRealUserId"`          // 关联wei_xin_real_user表的id
	Price                 float64 `gorm:"price" json:"price"`                                    // 交易单价
	Status                int     `gorm:"status" json:"status"`                                  // 订单票务状态,0表示未支付,1表示已支付,2表示支付失败,3表示取消订单,4表示退款(已经退款),5表示已经核销, 6表示在支付等待支付回调,7表示申请退款(还没退钱),8:拒绝申请退款, 9:退款失败
	Remark                string  `gorm:"remark" json:"remark"`                                  // 备注信息
	OrderId               int     `gorm:"order_id" json:"orderId"`                               // 关联的订单id
	VerificationId        int     `gorm:"verification_id" json:"verificationId"`                 // 核销员ID
	VerificationAt        string  `gorm:"verification_at" json:"verificationAt"`                 // 核销时间
	RealUserAddress       string  `gorm:"real_user_address" json:"realUserAddress"`              // 实名的性别
	ActivityEndTime       string  `gorm:"activity_end_time" json:"activityEndTime"`              // 活动结束时间
	RealUserNickName      string  `gorm:"real_user_nick_name" json:"realUserNickName"`           // 实名的昵称
	RealUserMobile        string  `gorm:"real_user_mobile" json:"realUserMobile"`                // 实名的手机号码
	RealUserGender        string  `gorm:"real_user_gender" json:"realUserGender"`                // 实名的性别
	RealUserAge           string  `gorm:"real_user_age" json:"realUserAge"`                      // 实名的性别
	IsVerification        int     `gorm:"is_verification" json:"isVerification"`                 // 是否核销,0表示未核销,1表示已核销,2表示票品过期
	RealSerTeamName       string  `gorm:"real_ser_team_name" json:"realSerTeamName"`             // 战队名称
	ActivityStartTime     string  `gorm:"activity_start_time" json:"activityStartTime"`          // 活动开始时间
	ApplyStartDateTime    string  `gorm:"apply_start_date_time" json:"applyStartDateTime"`       // 申请退票时间
	ConfirmRefundDateTime string  `gorm:"confirm_refund_date_time" json:"confirmRefundDateTime"` // 确认退款时间
	CloseRefundDateTime   string  `gorm:"close_refund_date_time" json:"closeRefundDateTime"`     // 退款关闭时间
	FailRefundDateTime    string  `gorm:"fail_refund_date_time" json:"failRefundDateTime"`       // 退款失败时间
	RefundNo              string  `gorm:"refund_no" json:"refundNo"`                             // 自定义退款单号
	RefundReason          string  `gorm:"refund_reason" json:"refundReason"`                     // 申请退款原因
	RefundChannel         string  `gorm:"refund_channel" json:"refundChannel"`                   // 退款渠道
	TransactionNo         string  `gorm:"transaction_no" json:"transactionNo"`                   // 退款第三方交易号
	RejectRefundReason    string  `gorm:"reject_refund_reason" json:"rejectRefundReason"`        // 拒绝退款原因
	Stage                 int     `gorm:"stage" json:"stage"`                                    // 订单阶段(0表待结算[可以退票]，1表示可以结算)
	CardNo                string  `gorm:"card_no" json:"cardNo"`                                 // 身份证号码
	IdCardName            string  `gorm:"id_card_name" json:"idCardName"`                        // 身份证上的姓名
}

func (*OrderListEntity) TableName() string {
	return "order_list"
}
