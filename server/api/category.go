package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"wj/constant"
	"wj/models/app"
	"wj/response"
	"wj/service"
)

type AppCategory struct {
	service.AppCategoryService
}

func GetAppCategory() *AppCategory {
	return &AppCategory{}
}

func (c *AppCategory) GetCategoryList(context *gin.Context) {
	var param app.CategoryQueryParam
	sid, exists := context.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	categoryList := c.GetList(param)
	response.Success(constant.Selected, categoryList, context)
}
