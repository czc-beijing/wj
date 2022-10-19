package service

import (
	"wj/constant"
	"wj/global"
	"wj/models/app"
)

type AppCategoryService struct {
}

// 获取服务类目列表
func (g *AppCategoryService) GetList(param app.CategoryQueryParam) []app.Category {
	var categoryList []app.Category
	query := &app.Category{
		Sid:    param.Sid,
		Status: constant.OnLineStatus,
	}
	global.Db.Table("category").Where(query).Find(&categoryList)
	return categoryList
}
