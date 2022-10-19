package app

import "wj/models"

// 用户数据映射模型
type User struct {
	Id       uint64 `gorm:"primaryKey"` // 用户编号
	Avatar   string `gorm:"avatar"`     // 用户编号
	Nickname string `gorm:"nickname"`   // 用户编号
	OpenId   string `gorm:"open_id"`    // 微信用户唯一标识
	Username string `gorm:"username"`   // 用户名称
	Password string `gorm:"password"`   // 用户密码
	Sid      int64  `gorm:"sid"`        //
	Status   uint   `gorm:"status"`     // 用户状态
	Created  string `gorm:"created"`    // 创建时间
	Updated  string `gorm:"updated"`    // 更新时间
}

// 用户登录凭证校验模型
type Code2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 凭证校验后返回的JSON数据包模型
type Code2SessionResult struct {
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 用户信息,OpenID用户唯一标识
type UserInfo struct {
	Id       uint64 `json:"id" json:"id"`
	OpenId   string `json:"openId" json:"openId"`
	Token    string `json:"token" json:"token"`
	Nickname string `form:"nickname" json:"nickname"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// 用户登录参数模型
type UserInfoParam struct {
	OpenId   string `form:"openId"`
	Nickname string `form:"nickname"`
	Avatar   string `form:"avatar"`
	Sid      int64  `form:"sid"`
}

type UserServiceQueryParam struct {
	Page       models.AppPage
	CategoryId uint64 `form:"category_id"`
	Type       int    `form:"type" `
	Sid        uint64 `form:"sid"`
	OpenId     string `form:"openId"`
}

// 统计当日数据传输模型
type TodayDate struct {
	Pending     int64 `json:"pending"`     // 待审核
	Unpublished int64 `json:"unpublished"` // 未发布
	Published   int64 `json:"published"`   // 发布
	OffShelves  int64 `json:"off_shelves"` // 已下架服务
}

type OrderTodayDate struct {
	Unapproved  int64 `json:"unapproved"`  // 待同意
	Unpaid      int64 `json:"unpaid"`      // 待支付
	Unconfirmed int64 `json:"unconfirmed"` // 待确认
	Unrated     int64 `json:"unrated"`     // 待评价
}
