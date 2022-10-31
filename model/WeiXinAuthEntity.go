package model

import "time"

type WeiXinAuthEntity struct {
	Id        int64     `gorm:"id" json:"id"`                // 主键id
	IsDel     int64     `gorm:"is_del" json:"isDel"`         // 是否删除,1表示删除,0表示正常
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"` // 更新时间
	Type      int64     `gorm:"type" json:"type"`            // 类型
	OpenId    string    `gorm:"open_id" json:"openId"`       // open_id
	UserId    string    `gorm:"user_id" json:"userId"`       // user_id
	DeletedAt time.Time `gorm:"deleted_at" json:"deletedAt"` // 软删除时间
}

func (t *WeiXinAuthEntity) TableName() string {
	return "wei_xin_auth"
}
