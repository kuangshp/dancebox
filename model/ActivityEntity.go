package model

type ActivityEntity struct {
	BaseEntity
	Title          string    `gorm:"column:title" json:"title"`                             // 活动标题
	Status         int64     `gorm:"column:status" json:"status"`                           // 上下架，1表示上架，0表示下架
	TypeID         int       `gorm:"column:type_id" json:"typeId"`                          // 活动类型Id
	LevelID        int       `gorm:"column:level_id" json:"levelId"`                        // 活动级别Id
	SponsorID      int       `gorm:"column:sponsor_id" json:"sponsorId"`                    // 主办方id
	ProvinceID     int       `gorm:"column:province_id" json:"provinceId"`                  // 省份Id
	CityID         int       `gorm:"column:city_id" json:"cityId"`                          // 城市Id
	Address        string    `gorm:"column:address" json:"address"`                         // 具体地址
	Location       string    `gorm:"column:location" json:"location"`                       // 具体地址经纬度
	StartTime      LocalTime `gorm:"column:start_time" json:"startTime"`                    // 开始时间
	EndTime        LocalTime `gorm:"column:end_time" json:"endTime"`                        // 结束时间
	Cover          string    `gorm:"column:cover" json:"cover"`                             // 封面图
	Dances         string    `gorm:"column:dances" json:"dances"`                           // 存放舞种的id数组字符串
	Content        string    `gorm:"column:content" json:"content"`                         // 活动具体内容
	Stage          int       `gorm:"column:stage;default:0;NOT NULL" json:"stage"`          // 阶段,0表示待审核,1表示审核通过,2表示拒绝
	Reason         string    `gorm:"column:reason" json:"reason"`                           // 拒绝原因
	Remark         string    `gorm:"column:remark" json:"remark"`                           // 拒绝原因备注
	IsTop          int       `gorm:"column:is_top;default:0" json:"isTop"`                  // 是否置顶,0是非置顶，1是置顶
	ActivityQrCode string    `gorm:"column:activity_qr_code" json:"activityQrCode"`         // 小程序对应的活动详情二维码
	FullAddress    string    `gorm:"column:full_address" json:"fullAddress"`                // 地址全称
	DancesStr      string    `gorm:"column:dances_str" json:"dancesStr"`                    // 舞蹈中文名字列表，多个用英文,拼接
	IsTicket       int       `gorm:"column:is_ticket;default:0;NOT NULL" json:"isTicket"`   // 是否开通票务
	LookCount      int       `gorm:"column:look_count;default:0" json:"lookCount"`          // 浏览数
	IDCardChecked  int       `gorm:"column:id_card_checked;default:0" json:"IDCardChecked"` // 是否需要验证身份证
	IsRefund       int       `gorm:"column:is_refund;default:0" json:"isRefund"`            // 是否开始退款,0表示不开启,1表示开启
}

func (a *ActivityEntity) TableName() string {
	return "activity"
}
