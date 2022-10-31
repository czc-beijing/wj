package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"wj/constant"
	"wj/models/app"
	"wj/response"
	"wj/service"
)

type AppOrder struct {
	service.AppOrderService
}

func GetAppOrder() *AppOrder {
	return &AppOrder{}
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
	if insertId := o.Create(context, param); insertId > 0 {
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

func (o *AppOrder) GetUserOrderList(c *gin.Context) {
	var param app.MyOrderQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	orderList, rows := o.GetListByOpenId(c, param)
	response.SuccessAppPage(constant.Selected, orderList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (o *AppOrder) GetUserOrderTodayDate(c *gin.Context) {
	var param app.MyOrderTodayParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	data := o.TodayData(c, param)
	response.Success(constant.LoginSuccess, data, c)
}
