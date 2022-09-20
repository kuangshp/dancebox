package model

type WeiXinUserEntity struct {
	BaseEntity
	GameName         string `gorm:"game_name" json:"gameName"`                  // 比赛名称
	Introduction     string `gorm:"introduction" json:"introduction"`           // 个人简介
	Birthday         string `gorm:"birthday" json:"birthday"`                   // 生日
	Constellation    string `gorm:"constellation" json:"constellation"`         // 星座
	DanceAge         string `gorm:"dance_age" json:"danceAge"`                  // 舞龄
	TeachingAge      string `gorm:"teaching_age" json:"teachingAge"`            // 教龄
	DanceListStr     string `gorm:"dance_list_str" json:"danceListStr"`         // 擅长的舞蹈列表字符串
	Mobile           string `gorm:"mobile" json:"mobile"`                       // 手机号码
	Gender           int    `gorm:"gender" json:"gender"`                       // 性别
	Country          string `gorm:"country" json:"country"`                     // 国家
	Province         string `gorm:"province" json:"province"`                   // 省份
	City             string `gorm:"city" json:"city"`                           // 城市
	Language         string `gorm:"language" json:"language"`                   // 语言
	Avatar           string `gorm:"avatar" json:"avatar"`                       // 头像
	Status           int    `gorm:"status" json:"status"`                       // 状态,0不正常,1正常
	LastLoginTime    string `gorm:"last_login_time" json:"lastLoginTime"`       // 最后登录时间
	OpenId           string `gorm:"open_id" json:"openId"`                      // 微信小程序授权的openId
	SessionKey       string `gorm:"session_key" json:"sessionKey"`              // 微信小程序授权的sessionKey
	IsRead           int    `gorm:"is_read" json:"isRead"`                      // 是否阅读购票协议,0表示未阅读,1表示已经阅读
	CountryCode      string `gorm:"country_code" json:"countryCode"`            // 手机号码所在地区(国家)
	NickName         string `gorm:"nick_name" json:"nickName"`                  // 昵称
	AvatarBanner     string `gorm:"avatar_banner" json:"avatarBanner"`          // 自动轮播头像
	RoleTags         string `gorm:"role_tags" json:"roleTags"`                  // 角色标签
	StyleDomain      string `gorm:"style_domain" json:"styleDomain"`            // 风格领域
	ContactList      string `gorm:"contact_list" json:"contactList"`            // 社交媒体
	DanceHistoryList string `gorm:"dance_history_list" json:"danceHistoryList"` // 街舞生涯
}

func (*WeiXinUserEntity) TableName() string {
	return "wei_xin_user"
}
