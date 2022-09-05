package model

type VerificationAccountEntity struct {
	BaseEntity
	NickName            string `gorm:"nick_name" json:"nickName"`                        // 昵称
	CountryCode         string `gorm:"country_code" json:"countryCode"`                  // 手机号码所在地区(国家)
	Mobile              string `gorm:"mobile" json:"mobile"`                             // 微信授权手机号码
	OpenId              string `gorm:"open_id" json:"openId"`                            // 微信小程序授权的openId
	SessionKey          string `gorm:"session_key" json:"sessionKey"`                    // 微信小程序授权的sessionKey
	VerificationAdminId int    `gorm:"verification_admin_id" json:"verificationAdminId"` // 关联到后台创建核销账号id
}

func (*VerificationAccountEntity) TableName() string {
	return "verification_account"
}
