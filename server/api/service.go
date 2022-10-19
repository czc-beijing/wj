package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"wj/constant"
	"wj/models/app"
	"wj/response"
	"wj/service"
)

type AppService struct {
	service.AppServiceService
}

func GetAppService() *AppService {
	return &AppService{}
}

func (s *AppService) GetServiceList(c *gin.Context) {
	var param app.ServiceQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	serviceList, rows := s.GetList(c, param)
	response.SuccessAppPage(constant.Selected, serviceList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (s *AppService) GetServiceInfo(c *gin.Context) {
	id := c.Param("id")
	serviceInfo := s.GetInfo(cast.ToUint64(id))
	response.Success(constant.Selected, serviceInfo, c)
}

func (s *AppService) CreateService(c *gin.Context) {
	var param app.ServiceCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	res := gin.H{
		"service_id": 0,
	}
	if insertId := s.Create(c, param); insertId > 0 {
		res["service_id"] = insertId
		response.Success(constant.Created, res, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}

func (s *AppService) UpdateServiceStatus(c *gin.Context) {
	var param app.ServiceStatusUpdateParam
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
	if count := s.UpdateStatus(param); count > 0 {
		res["result"] = true
		response.Success(constant.Updated, res, c)
		return
	}
	response.Failed(constant.NotUpdated, c)
}

func (s *AppService) UpdateService(c *gin.Context) {
	var param app.ServiceUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	res := gin.H{
		"service_id": 0,
	}
	if updateId := s.Update(c, param); updateId > 0 {
		res["service_id"] = updateId
		response.Success(constant.Updated, res, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}

func (s *AppService) MyServiceList(c *gin.Context) {
	var param app.UserServiceQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	ServiceList, rows := s.GetListByOpenId(c, param.Type, param.Page.PageNum, param.Page.PageSize)
	response.SuccessAppPage(constant.Selected, ServiceList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (s *AppService) MyServiceTodayDate(c *gin.Context) {
	var param app.UserServiceQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	data := s.TodayData(c, param)
	response.Success(constant.LoginSuccess, data, c)
}
