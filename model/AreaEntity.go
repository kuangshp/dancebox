package model

type AreaEntity struct {
	ID         int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键id
	Pid        int    `json:"pid" gorm:"column:pid;default:0;NOT NULL"`       // 父节点id
	Shortname  string `json:"shortname" gorm:"column:shortname;NOT NULL"`     // 短名字
	Name       string `json:"name" gorm:"column:name;NOT NULL"`               // 名称
	MergerName string `json:"mergerName" gorm:"column:merger_name"`           // 全称
	Level      int    `json:"level" gorm:"column:level"`                      // 级别
	Pinyin     string `json:"pinyin" gorm:"column:pinyin"`                    // 拼音写法
	Code       string `json:"code" gorm:"column:code"`                        // code
	ZipCode    string `json:"zipCode" gorm:"column:zip_code"`                 // 邮编
	First      string `json:"first" gorm:"column:first"`                      // first
	Lng        string `json:"lng" gorm:"column:lng"`                          // 经度
	Lat        string `json:"lat" gorm:"column:lat"`                          // 维度
}

func (m *AreaEntity) TableName() string {
	return "area"
}
