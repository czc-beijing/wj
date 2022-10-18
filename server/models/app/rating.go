package app

import "imall/models"

// 服务类目映射模型
type Rating struct {
	Id           uint64 `gorm:"primaryKey"`   // 类目编号
	OpenId       string `gorm:"open_id"`      // 类目编号
	OrderId      uint64 `gorm:"order_id"`     // 类目编号
	ServiceId    uint64 `gorm:"service_id"`   // 类目编号
	Score        int64  `gorm:"score"`        // 类目名称
	Content      string `gorm:"content"`      // 父级编号
	Illustration string `gorm:"illustration"` // 类目级别
	Created      string `gorm:"created"`      // 创建时间
	Updated      string `gorm:"updated"`      // 更新时间
	Sid          uint64 `gorm:"sid"`          // 店铺编号
}

// 评论列表参数模型
type RatingQueryParam struct {
	Page    models.AppPage
	OrderId uint64 `form:"order_id"`
	Sid     uint64 `form:"sid"`
	OpenId  string `form:"openId"`
}

// 评论列表参数模型
type RatingListParam struct {
	Page      models.AppPage
	ServiceId uint64 `form:"service_id"`
	Sid       uint64 `form:"sid"`
	OpenId    string `form:"openId"`
}

// 评论搜索参数模型
type RatingSearchParam struct {
	KeyWord string `form:"keyWord" binding:"required"`
}

// 评论列表传输模型
type RatingInfo struct {
	ID           uint64   `json:"id"`
	Score        int64    `json:"score"`
	Content      string   `json:"content"`
	Illustration []string `json:"illustration"`
	CreateTime   string   `json:"create_time"`
	Author       struct {
		ID       uint64 `json:"id"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	} `json:"author"`
}

type RatingCreateParam struct {
	OrderId      uint64   `from:"order_id" json:"order_id"`
	Score        int64    `from:"score" json:"score"`
	Content      string   `form:"content" json:"content"`
	Illustration []string `form:"illustration" json:"illustration"`
	Sid          uint64   `form:"sid" json:"sid"`
	OpenId       string   `form:"openId" json:"openId"`
}
