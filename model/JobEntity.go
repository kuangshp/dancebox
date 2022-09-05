package model

type JobEntity struct {
	BaseEntity
	JobName       string `gorm:"job_name" json:"jobName"`              // 职位名称
	JobType       int64  `gorm:"job_type" json:"jobType"`              // 职位类型(0: 全职,1:兼职)
	SalaryType    int64  `gorm:"salary_type" json:"salaryType"`        // 招聘薪资类型(0: 面议,1:非面议)
	SalaryOne     string `gorm:"salary_one" json:"salaryOne"`          // 兼职的时候每次薪资
	MinSalary     string `gorm:"min_salary" json:"minSalary"`          // 最低薪资
	MaxSalary     string `gorm:"max_salary" json:"maxSalary"`          // 最高薪资
	ProvinceId    int64  `gorm:"province_id" json:"provinceId"`        // 省份Id
	CityId        int64  `gorm:"city_id" json:"cityId"`                // 城市Id
	Address       string `gorm:"address" json:"address"`               // 具体地址
	FullAddress   string `gorm:"full_address" json:"fullAddress"`      // 地址全称
	Location      string `gorm:"location" json:"location"`             // 具体地址经纬度
	Content       string `gorm:"content" json:"content"`               // 职位描述
	IsTop         int64  `gorm:"is_top" json:"isTop"`                  // 是否置顶,0是非置顶，1是置顶
	Status        int64  `gorm:"status" json:"status"`                 // 1:表示上架,2:表示下架
	ResumeCount   int64  `gorm:"resume_count" json:"resumeCount"`      // 简历数量
	JobQrCode     string `gorm:"job_qr_code" json:"jobQrCode"`         // 小程序对应的招聘二维码
	SponsorId     int64  `gorm:"sponsor_id" json:"sponsorId"`          // 关联主办方ID
	SalaryOneType int64  `gorm:"salary_one_type" json:"salaryOneType"` // 兼职薪资类型(0:元/天,1:元/次)
}

func (t *JobEntity) TableName() string {
	return "job"
}
