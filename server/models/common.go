package models

// Page 分页参数模型
type Page struct {
	PageNum  int `form:"page"  json:"page"`
	PageSize int `form:"count" json:"count"`
}

// AppPage 分页参数模型
type AppPage struct {
	PageNum  int `form:"page"  json:"page"`
	PageSize int `form:"count" json:"count"`
}
