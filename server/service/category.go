package service

import (
	"imall/common"
	"imall/constant"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
)

type WebCategoryService struct {
}

type AppCategoryService struct {
}

// 创建服务类目
func (c *WebCategoryService) Create(param web.CategoryCreateParam) int64 {
	category := web.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
		Sid:      param.Sid,
	}
	return global.Db.Create(&category).RowsAffected
}

// 删除服务类目
func (c *WebCategoryService) Delete(param web.CategoryDeleteParam) int64 {
	rows := global.Db.Where("parent_id = ?", param.Id).Delete(&web.Category{}).RowsAffected
	if rows < 0 {
		return 0
	}
	return global.Db.Delete(&web.Category{}, param.Id).RowsAffected
}

// 更新服务类目
func (c *WebCategoryService) Update(param web.CategoryUpdateParam) int64 {
	category := web.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(&category).RowsAffected
}

// 获取服务类目列表
func (g *WebCategoryService) GetList(param web.CategoryQueryParam) []web.CategoryList {
	var categoryList []web.CategoryList
	query := &web.Category{
		Sid:    param.Sid,
		Status: constant.OnLineStatus,
	}
	global.Db.Table("category").Where(query).Find(&categoryList)
	return categoryList
}

// 获取服务类目选项
func (c *AppCategoryService) GetOption(param app.CategoryQueryParam) (option []app.CategoryOption) {
	categorys := make([]app.Category, 0)
	categoryOptions := make([]app.CategoryOption, 0)
	global.Db.Table("category").Where("level = 1 and sid =?", param.Sid).Find(&categorys)
	for _, item := range categorys {
		option := app.CategoryOption{
			Id:   item.Id,
			Text: item.Name,
		}
		categoryOptions = append(categoryOptions, option)
	}
	return categoryOptions
}
