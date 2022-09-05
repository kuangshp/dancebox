package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"dancebox-admin-api/global"
	"dancebox-admin-api/utils"
)

type SponsorEntity struct {
	BaseEntity
	Name           string        `gorm:"column:name" json:"name"`                                          // 主办方名称
	Avatar         string        `gorm:"column:avatar" json:"avatar"`                                      // 头像
	ProvinceID     int64         `gorm:"column:province_id" json:"provinceId"`                             // 省份Id
	CityID         int64         `gorm:"column:city_id" json:"cityId"`                                     // 城市Id
	Address        string        `gorm:"column:address" json:"address"`                                    // 具体地址
	UserName       string        `gorm:"column:user_name" json:"userName"`                                 // 用户名
	Mobile         string        `gorm:"column:mobile" json:"mobile"`                                      // 手机号码
	Email          string        `gorm:"column:email" json:"email"`                                        // 邮箱
	HomeQr         string        `gorm:"column:home_qr" json:"homeQr"`                                     // 主办方小程序主页
	BannerUrls     LocalList     `gorm:"column:banner_urls" json:"bannerUrls"`                             // 轮播图图片
	Tags           LocalList     `gorm:"column:tags" json:"tags"`                                          // 主办方标签
	FullAddress    string        `gorm:"column:full_address" json:"fullAddress"`                           // 地址全称
	Location       string        `gorm:"column:location" json:"location"`                                  // 具体地址经纬度
	Description    string        `gorm:"column:description" json:"description"`                            // 描素
	IsTop          sql.NullInt64 `gorm:"column:is_top;default:0" json:"isTop"`                             // 是否置顶,0是非置顶，1是置顶
	WechatUrl      *string       `gorm:"column:wechat_url" json:"wechatUrl"`                               // 微信二维码
	WechatName     *string       `gorm:"column:wechat_name" json:"wechatName"`                             // 微信号
	LookCount      int           `gorm:"column:look_count;default:0" json:"lookCount"`                     // 浏览数
	Openid         string        `gorm:"column:openid" json:"openid"`                                      // 订阅号的openid
	Unionid        string        `gorm:"column:unionid" json:"unionid"`                                    // 订阅号的unionid
	Language       string        `gorm:"column:language" json:"language"`                                  // language,zh_CN 简体，zh_TW 繁体，en 英语
	IDCheckedTotal int           `gorm:"column:id_checked_total;default:0;NOT NULL" json:"IDCheckedTotal"` // 验证身份证可调用总条数
	IDCheckedCount int           `gorm:"column:id_checked_count;default:0;NOT NULL" json:"IDCheckedCount"` // 验证身份证可调用条数
	IsCloseInform  int           `gorm:"column:is_close_inform;default:0;NOT NULL" json:"isCloseInform"`   // 是否关闭后台提示,0表示不关闭,1表示关闭
	IsLimitInform  int           `gorm:"column:is_limit_inform;default:0;NOT NULL" json:"isLimitInform"`   // 是否已发送身份验证短信调用次数不足,0表示还未发送，1表示已经发送过
}

func (s *SponsorEntity) TableName() string {
	return "sponsor"
}

// BeforeCreate 创建前的钩子函数
func (s *SponsorEntity) BeforeCreate(tx *gorm.DB) (err error) {
	return getAddressByCityId(s, tx)
}

// BeforeUpdate 更新数据的钩子函数
func (s *SponsorEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	return getAddressByCityId(s, tx)
}

// getAddressByCityId 根据城市id获取具体地址及经纬度
func getAddressByCityId(s *SponsorEntity, tx *gorm.DB) (err error) {
	// 根据城市id查询地址信息
	var areaEntity AreaEntity
	if result := tx.Where("id=?", s.CityID).First(&areaEntity).Error; result != nil {
		global.Logger.Error("主办方表的钩子函数中查询地址失败" + result.Error())
		return errors.New("创建错误")
	}
	var addressList = strings.Split(areaEntity.MergerName, ",")
	// 去除第一个元素
	addressList = addressList[1:]
	var address = strings.Join(addressList, "") + s.Address
	s.FullAddress = address
	// 根据地址查询经纬度
	lng, lat := utils.GetLngLatByAddress(address)
	s.Location = fmt.Sprintf("%f,%f", lat, lng)
	return
}
