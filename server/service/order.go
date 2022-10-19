package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
	"wj/common"
	"wj/constant"
	"wj/global"
	"wj/models/app"
	"wj/utils"
)

type AppOrderService struct {
}

func NewAppOrderService() *AppOrderService {
	return &AppOrderService{}
}

// 创建服务
func (o *AppOrderService) Create(c *gin.Context, param app.OrderCreateParam) uint64 {

	//查询服务信息
	consumerOpenId, _ := c.Get("openId")
	sid, _ := c.Get("sid")
	//查询发布者信息
	var ServiceInfo app.Service
	global.Db.Table("Service").Where(map[string]interface{}{"id": param.ServiceID}).First(&ServiceInfo)
	if ServiceInfo.Id <= 0 {
		return 0
	}
	ServiceData := NewAppServiceService().GetInfo(ServiceInfo.Id)
	serviceSnap, _ := json.Marshal(ServiceData)
	orderNo := utils.GenerateOrderSn(ServiceInfo.Id)

	order := app.Order{
		OrderNo:          orderNo,
		PublisherOpenId:  ServiceInfo.OpenId,
		ConsumerOpenId:   consumerOpenId.(string),
		ServiceID:        param.ServiceID,
		NationalCodeFull: param.Address.NationalCodeFull,
		TelNumber:        param.Address.TelNumber,
		UserName:         param.Address.UserName,
		NationalCode:     param.Address.NationalCode,
		PostalCode:       param.Address.PostalCode,
		ProvinceName:     param.Address.ProvinceName,
		CityName:         param.Address.CityName,
		CountyName:       param.Address.CountyName,
		StreetName:       param.Address.StreetName,
		DetailInfoNew:    param.Address.DetailInfoNew,
		DetailInfo:       param.Address.DetailInfo,
		BeginDate:        param.BeginDate,
		EndDate:          param.EndDate,
		ServiceSnap:      cast.ToString(serviceSnap),
		Sid:              cast.ToUint64(sid),
		Description:      param.Description,
		Status:           0,
		Created:          time.Now().Format("20060102"),
		Updated:          time.Now().Format("20060102"),
	}
	_ = global.Db.Create(&order).RowsAffected
	return order.Id
}

func (o *AppOrderService) GetListByOpenId(c *gin.Context, param app.MyOrderQueryParam) ([]app.OrderInfo, int64) {
	sid, _ := c.Get("sid")
	param.Sid = cast.ToUint64(sid)
	openID, _ := c.Get("openId")
	param.OpenId = openID.(string)

	var orderList []app.OrderInfo
	var orders []app.Order
	query := map[string]interface{}{"sid": param.Sid}
	if len(param.Status) > 0 {
		query["status"] = param.Status
	}
	if param.Role == 1 { //1：发布者
		query["publisher_open_id"] = param.OpenId
	} else if param.Role == 2 {
		query["consumer_open_id"] = param.OpenId
	}
	var rows int64
	if param.Page.PageNum > 0 && param.Page.PageSize > 0 {
		offset := (param.Page.PageNum - 1) * param.Page.PageSize
		listDao := global.Db.Offset(offset).Limit(param.Page.PageSize).Table("order").Where(query).Group("id desc")
		rowDao := global.Db.Table("order").Where(query)
		if len(param.UserName) > 0 {
			listDao.Where("user_name LIKE ?", "%"+param.UserName+"%")
			rowDao.Where("user_name LIKE ?", "%"+param.UserName+"%")
		}
		rows = rowDao.Find(&[]app.Service{}).RowsAffected
		listDao.Find(&orders)
	}
	if len(orders) <= 0 {
		return []app.OrderInfo{}, 0
	}
	var openIDs []string
	for _, item := range orders {
		openIDs = append(openIDs, item.PublisherOpenId)
		openIDs = append(openIDs, item.ConsumerOpenId)
	}
	var userList []app.User
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	userMap := map[string]app.User{}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	for _, item := range orders {
		var ServiceInfo app.ServiceInfo
		_ = json.Unmarshal([]byte(item.ServiceSnap), &ServiceInfo)
		orderInfo := app.OrderInfo{
			ID:          item.Id,
			OrderNo:     item.OrderNo,
			Price:       ServiceInfo.Price,
			ServiceSnap: ServiceInfo,
			AddressSnap: app.AddressInfo{
				CityName:     item.CityName,
				UserName:     item.UserName,
				TelNumber:    item.TelNumber,
				CountyName:   item.CountyName,
				DetailInfo:   item.DetailInfo,
				PostalCode:   item.PostalCode,
				NationalCode: item.NationalCode,
				ProvinceName: item.ProvinceName,
			},
			Status:     item.Status,
			CreateTime: item.Created,
		}
		orderInfo.Role = 2
		if param.OpenId == item.PublisherOpenId {
			orderInfo.Role = 1
		}
		orderInfo.Description = item.Description
		if userInfo, exists := userMap[item.PublisherOpenId]; exists {
			orderInfo.Publisher.ID = userInfo.Id
			orderInfo.Publisher.Nickname = userInfo.Nickname
			orderInfo.Publisher.Avatar = userInfo.Avatar
			orderInfo.Publisher.Tel = "13429208394"
			orderInfo.PublisherTel = "13429208394"
		}
		orderInfo.BeginDate = item.BeginDate
		orderInfo.EndDate = item.EndDate
		orderInfo.PublisherTel = "13429208394"
		if userInfo, exists := userMap[item.ConsumerOpenId]; exists {
			orderInfo.Consumer.ID = userInfo.Id
			orderInfo.Consumer.Nickname = userInfo.Nickname
			orderInfo.Consumer.Avatar = userInfo.Avatar
			orderInfo.Consumer.Tel = item.TelNumber
		}
		orderInfo.Tel = item.TelNumber
		orderList = append(orderList, orderInfo)
	}
	return orderList, rows
}

func (o *AppOrderService) GetDetail(param app.OrderDetailParam) app.OrderInfo {
	var orderData *app.Order
	_ = global.Db.Table("order").Where("id = ?", param.Id).First(&orderData).RowsAffected
	if orderData.Id == 0 {
		return app.OrderInfo{}
	}
	var openIDs []string
	openIDs = append(openIDs, orderData.PublisherOpenId)
	openIDs = append(openIDs, orderData.ConsumerOpenId)
	var userList []app.User
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	userMap := map[string]app.User{}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	var ServiceInfo app.ServiceInfo
	_ = json.Unmarshal([]byte(orderData.ServiceSnap), &ServiceInfo)
	orderInfo := app.OrderInfo{
		ID:          orderData.Id,
		OrderNo:     orderData.OrderNo,
		Price:       ServiceInfo.Price,
		ServiceSnap: ServiceInfo,
		AddressSnap: app.AddressInfo{
			CityName:     orderData.CityName,
			UserName:     orderData.UserName,
			TelNumber:    orderData.TelNumber,
			CountyName:   orderData.CountyName,
			DetailInfo:   orderData.DetailInfo,
			PostalCode:   orderData.PostalCode,
			NationalCode: orderData.NationalCode,
			ProvinceName: orderData.ProvinceName,
		},
		Status:     orderData.Status,
		CreateTime: orderData.Created,
	}
	orderInfo.Role = 2
	if param.OpenId == orderData.PublisherOpenId {
		orderInfo.Role = 1
	}
	if userInfo, exists := userMap[orderData.PublisherOpenId]; exists {
		orderInfo.Publisher.ID = userInfo.Id
		orderInfo.Publisher.Nickname = userInfo.Nickname
		orderInfo.Publisher.Avatar = userInfo.Avatar
		orderInfo.Publisher.Tel = "13429208394"
		orderInfo.PublisherTel = "13429208394"
	}
	if userInfo, exists := userMap[orderData.ConsumerOpenId]; exists {
		orderInfo.Consumer.ID = userInfo.Id
		orderInfo.Consumer.Nickname = userInfo.Nickname
		orderInfo.Consumer.Avatar = userInfo.Avatar
		orderInfo.Consumer.Tel = orderData.TelNumber
	}
	orderInfo.Tel = orderData.TelNumber
	return orderInfo
}

// 统计当日数据
func (o *AppOrderService) TodayData(c *gin.Context, param app.MyOrderTodayParam) app.OrderTodayDate {
	sid, _ := c.Get("sid")
	openID, _ := c.Get("openId")
	param.Sid = cast.ToUint64(sid)
	param.OpenId = openID.(string)

	pps := "consumer_open_id = ? and sid = ? and status = 0"
	pds := "consumer_open_id = ? and sid = ? and status = 1"
	prs := "consumer_open_id = ? and sid = ? and status = 2"
	ras := "consumer_open_id = ? and sid = ? and status = 3"
	if param.Role == 1 {
		pps = "publisher_open_id = ? and sid = ? and status = 0"
		pds = "publisher_open_id = ? and sid = ? and status = 1"
		prs = "publisher_open_id = ? and sid = ? and status = 2"
		ras = "publisher_open_id = ? and sid = ? and status = 3"
	}
	orderDate := app.OrderTodayDate{}
	global.Db.Table("order").Where(pps, param.OpenId, param.Sid).Count(&orderDate.Unapproved)
	global.Db.Table("order").Where(pds, param.OpenId, param.Sid).Count(&orderDate.Unpaid)
	global.Db.Table("order").Where(prs, param.OpenId, param.Sid).Count(&orderDate.Unconfirmed)
	global.Db.Table("order").Where(ras, param.OpenId, param.Sid).Count(&orderDate.Unrated)
	return orderDate
}

// 更新订单
func (o *AppOrderService) UpdateStatus(param app.OrderUpdateParam) uint64 {
	updateStatus := constant.ActionStatus[param.Action]
	order := app.Order{
		Id:      param.Id,
		Status:  updateStatus,
		Updated: common.NowTime(),
	}
	_ = global.Db.Model(&order).Where(map[string]interface{}{"id": param.Id}).Updates(&order).RowsAffected
	return order.Id
}
