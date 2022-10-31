package model

type WechatComponentsDetailEntity struct {
	BaseEntity
	ImgUrl             string `gorm:"img_url" json:"imgUrl"`                          // 图片地址
	LinkType           int64  `gorm:"link_type" json:"linkType"`                      // 跳转类型
	WechatId           string `gorm:"wechat_id" json:"wechatId"`                      // 微信小程序appid
	WechatPath         string `gorm:"wechat_path" json:"wechatPath"`                  // 微信小程序路径
	LinkImgUrl         string `gorm:"link_img_url" json:"linkImgUrl"`                 // 跳转图片
	LinkText           string `gorm:"link_text" json:"linkText"`                      // 跳转文本
	WechatComponentsId int64  `gorm:"wechat_components_id" json:"wechatComponentsId"` // 关联到组件表的主键id
}

func (t *WechatComponentsDetailEntity) TableName() string {
	return "wechat_components_detail"
}
