package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"imall/constant"
	"imall/models/app"
	"imall/response"
	"imall/service"
)

type WebServices struct {
	service.WebServicesService
}

type AppServices struct {
	service.AppServicesService
}

func GetWebservices() *WebServices {
	return &WebServices{}
}

func GetAppServices() *AppServices {
	return &AppServices{}
}

func (g *AppServices) GetServicesList(c *gin.Context) {
	var param app.ServicesQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	servicesList, rows := g.GetList(param)
	response.SuccessAppPage(constant.Selected, servicesList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (g *AppServices) GetServicesInfo(c *gin.Context) {
	id := c.Param("id")
	servicesInfo := g.GetInfo(cast.ToUint64(id))
	response.Success(constant.Selected, servicesInfo, c)
}

func (g *AppServices) CreateServices(c *gin.Context) {
	var param app.ServicesCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	res := gin.H{
		"service_id": 0,
	}
	if insertId := g.Create(param); insertId > 0 {
		res["service_id"] = insertId
		response.Success(constant.Created, res, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}

func (g *AppServices) UpdateServicesStatus(c *gin.Context) {
	var param app.ServicesStatusUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	id, exists := c.Params.Get("id")
	if exists {
		param.Id = cast.ToUint64(id)
	}
	res := gin.H{
		"result": false,
	}
	if count := g.UpdateStatus(param); count > 0 {
		res["result"] = true
		response.Success(constant.Updated, res, c)
		return
	}
	response.Failed(constant.NotUpdated, c)
}

func (g *AppServices) UpdateServices(c *gin.Context) {
	var param app.ServicesUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	res := gin.H{
		"service_id": 0,
	}
	if updateId := g.Update(param); updateId > 0 {
		res["service_id"] = updateId
		response.Success(constant.Updated, res, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}
