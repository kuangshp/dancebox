package model

type VjThemeDetailEntity struct {
	BaseEntity
	Level   int64  `gorm:"level" json:"level"`      // 比赛级别：比如32强、16强、8强、4强、2强
	Url     string `gorm:"url" json:"url"`          // 主题背景图地址
	ThemeId int64  `gorm:"theme_id" json:"themeId"` // 关联vj_theme表主键id
}

func (t *VjThemeDetailEntity) TableName() string {
	return "vj_theme_detail"
}
