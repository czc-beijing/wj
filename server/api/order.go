package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"imall/constant"
	"imall/models/app"
	"imall/response"
	"imall/service"
)

type WebOrder struct {
	service.WebOrderService
}

type AppOrder struct {
	service.AppOrderService
}

func GetWebOrder() *WebOrder {
	return &WebOrder{}
}

func GetAppOrder() *AppOrder {
	return &AppOrder{}
}

func (o *AppOrder) GetOrderList(context *gin.Context) {
	var param app.OrderQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	orderList := o.GetList(param)
	response.Success("查询成功", orderList, context)
}

func (o *AppOrder) CreateOrder(context *gin.Context) {
	var param app.OrderCreateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	_, exists := context.Get("sid")
	if !exists {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	_, exists = context.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	res := gin.H{
		"service_id": 0,
	}
	if insertId := service.NewAppOrderService().Create(context, param); insertId > 0 {
		res["service_id"] = insertId
		response.Success(constant.Created, res, context)
		return
	}
	response.Failed(constant.NotCreated, context)
}

func (o *AppOrder) UpdateOrderStatus(context *gin.Context) {
	var param app.OrderUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	res := gin.H{
		"result": false,
	}
	param.Id = cast.ToUint64(context.Param("id"))
	if count := o.UpdateStatus(param); count > 0 {
		res["result"] = true
		response.Success("更新成功", res, context)
		return
	}
	response.FailedData("更新失败", res, context)
}

func (o *AppOrder) GetOrderDetail(c *gin.Context) {
	var param app.OrderDetailParam
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	openID, exists := c.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	param.Id = cast.ToUint64(c.Param("id"))
	param.OpenId = openID.(string)
	productDetail := o.GetDetail(param)
	response.Success("查询成功", productDetail, c)
}
