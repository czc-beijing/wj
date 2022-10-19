package app

import "wj/models"

// 服务映射模型
type Service struct {
	Id              uint64 `gorm:"id"`               // 服务编号
	OpenId          string `gorm:"open_id"`          // openId 发布者
	CategoryId      uint64 `gorm:"category_id"`      // 类目编号
	Type            int    `gorm:"type"`             // 服务类型
	Title           string `gorm:"title"`            // 服务标题
	DesignatedPlace int    `gorm:"designated_place"` // 预约类型
	Description     string `gorm:"description"`      // 描述
	Score           string `gorm:"score"`            // 描述
	Price           string `gorm:"price"`            // 服务价格
	Status          int    `gorm:"status"`           // 服务状态，1-出售中，2-仓库中
	CoverImage      string `gorm:"cover_image"`      // 服务图片
	SalesVolume     int    `gorm:"sales_volume"`     // 服务图片
	Tel             string `gorm:"tel"`              // 电话
	Sid             uint64 `gorm:"sid"`              // 店铺编号
	BeginDate       string `gorm:"begin_date"`       // 开始时间
	EndDate         string `gorm:"end_date"`         // 结束时间
	CreateTime      string `gorm:"create_time"`      // 创建时间
	UpdateTime      string `gorm:"update_time"`      // 更新时间
}

type ServiceQueryParam struct {
	Page       models.AppPage
	CategoryId uint64 `form:"category_id"`
	Type       uint64 `form:"type" `
	Sid        uint64 `form:"sid"`
}

type ServiceInfo struct {
	ID              uint64 `json:"id"`
	Type            int    `json:"type"`
	DesignatedPlace bool   `json:"designated_place"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Price           string `json:"price"`
	Score           string `json:"score"`
	SalesVolume     int    `json:"sales_volume"`
	CreateTime      string `json:"create_time"`
	Status          int    `json:"status"`
	BeginDate       string `json:"begin_date"`
	EndDate         string `json:"end_date"`
	Publisher       struct {
		ID       uint64 `json:"id"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		RealName string `json:"real_name"`
		Tel      string `json:"tel"`
	} `json:"publisher"`
	Category struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	CoverImage struct {
		ID   uint64 `json:"id"`
		Path string `json:"path"`
	} `json:"cover_image"`
}

type ServiceCreateParam struct {
	Type            int    `from:"type" json:"type"`
	Title           string `from:"title" json:"title"`
	OpenId          string `from:"open_id" json:"open_id"`
	CategoryId      uint64 `from:"category_id" binding:"required,gt=0" json:"category_id"`
	CoverImage      string `from:"cover_image" binding:"required" json:"cover_image"`
	DesignatedPlace bool   `from:"designated_place" json:"designated_place"`
	Description     string `from:"description" json:"description"`
	Price           string `from:"price" json:"price"`
	BeginDate       string `from:"begin_date" json:"begin_date"`
	EndDate         string `from:"end_date" json:"end_date"`
	Sid             uint64 `from:"sid" json:"sid"`
}

type ServiceStatusUpdateParam struct {
	Id     uint64 `json:"id" from:"id"`
	Action int    `json:"action" from:"action" binding:"required"`
}

type ServiceUpdateParam struct {
	Id                    uint64     `from:"id" json:"id"`
	Type                  int        `from:"type" json:"type"`
	Title                 string     `from:"title" json:"title"`
	OpenId                string     `from:"open_id" json:"open_id"`
	CategoryId            uint64     `from:"category_id"  json:"category_id"`
	CoverImage            CoverImage `from:"cover_image" json:"cover_image"`
	DesignatedPlace       bool       `from:"designated_place" json:"designated_place"`
	UpdateDesignatedPlace int        `from:"update_designated_place" json:"update_designated_place"`
	Description           string     `from:"description" json:"description"`
	Price                 string     `from:"price" json:"price"`
	BeginDate             string     `from:"begin_date" json:"begin_date"`
	EndDate               string     `from:"end_date" json:"end_date"`
	Sid                   uint64     `from:"sid" json:"sid"`
}

// 服务状态更新参数模型
type CoverImage struct {
	Id   uint64 `json:"id" from:"id"`
	Path string `json:"path" from:"path"`
}

type FileInfo struct {
	Id   uint64 `json:"id" from:"id"`
	Path string `json:"path" from:"path"`
}
