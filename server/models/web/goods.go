package web

import "imall/models"

// 服务映射模型
type Goods struct {
	Id              uint64 `gorm:"id"`               // 服务编号
	OpenId          string `gorm:"open_id"`          // openId 发布者
	CategoryId      uint64 `gorm:"category_id"`      // 类目编号
	Type            int    `gorm:"type"`             // 服务类型
	Title           string `gorm:"title"`            // 服务标题
	DesignatedPlace int    `gorm:"designated_place"` // 预约类型
	Description     string `gorm:"description"`      // 描述
	Score           string `gorm:"score"`            // 描述
	Price           string `gorm:"price"`            // 服务价格
	Status          uint   `gorm:"status"`           // 服务状态，1-出售中，2-仓库中
	ImageUrl        string `gorm:"image_url"`        // 服务图片
	SalesVolume     uint64 `gorm:"sales_volume"`     // 服务图片
	Tel             string `gorm:"tel"`              // 电话
	Sid             uint64 `gorm:"sid"`              // 店铺编号
	Created         string `gorm:"create_time"`      // 创建时间
	Updated         string `gorm:"update_time"`      // 更新时间
}

// 服务创建参数模型
type GoodsCreateParam struct {
	CategoryId uint64  `json:"categoryId" binding:"required,gt=0"`
	Title      string  `json:"title" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
	Quantity   uint    `json:"quantity" binding:"required,gt=0"`
	ImageUrl   string  `json:"imageUrl" binding:"required"`
	Remark     string  `json:"remark"`
	Sid        uint64  `json:"sid" binding:"required,gt=0"`
}

// 服务删除参数模型
type GoodsDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 服务更新参数模型
type GoodsUpdateParam struct {
	Id         uint64  `json:"id" binding:"required,gt=0"`
	CategoryId uint64  `json:"categoryId" binding:"required,gt=0"`
	Title      string  `json:"title" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
	Quantity   uint    `json:"quantity" binding:"required,gt=0"`
	ImageUrl   string  `json:"imageUrl" binding:"required"`
	Remark     string  `json:"remark"`
}

// 服务状态更新参数模型
type GoodsStatusUpdateParam struct {
	Id     uint64 `json:"id" binding:"required,gt=0"`
	Action uint   `json:"action" binding:"action,gt=0"`
}

// 服务列表查询参数模型
type GoodsListParam struct {
	Page       models.Page
	Id         uint64 `form:"id"`
	CategoryId uint64 `form:"categoryId"`
	Title      string `form:"title"`
	Status     uint   `form:"status"`
	Sid        uint64 `form:"sid" binding:"required,gt=0"`
}

// 服务列表传输模型
type GoodsList struct {
	Id         uint64  `json:"id"`
	CategoryId uint64  `json:"categoryId"`
	Title      string  `json:"title"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	ImageUrl   string  `json:"imageUrl"`
	Remark     string  `json:"remark"`
	Sales      uint    `json:"sales"`
	Status     uint    `json:"status"`
	Created    string  `json:"created"`
}
