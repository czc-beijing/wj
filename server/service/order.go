package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"imall/common"
	"imall/constant"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"imall/utils"
	"time"
)

type WebOrderService struct {
}

type AppOrderService struct {
}

func NewAppOrderService() *AppOrderService {
	return &AppOrderService{}
}

// 删除订单
func (o *WebOrderService) Delete(param web.OrderDeleteParam) int64 {
	return global.Db.Delete(&web.Order{}, param.Id).RowsAffected
}

// 更新订单
func (o *WebOrderService) Update(param web.OrderUpdateParam) int64 {
	order := web.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// 获取订单列表
func (o *WebOrderService) GetList(param web.OrderListParam) ([]web.OrderList, int64) {
	return nil, 0
}

// 更新订单
func (o *AppOrderService) UpdateStatus(param app.OrderUpdateParam) uint64 {
	updateStatus := constant.ActionStatus[param.Action]
	order := web.Order{
		Id:      param.Id,
		Status:  updateStatus,
		Updated: common.NowTime(),
	}
	_ = global.Db.Model(&order).Where(map[string]interface{}{"id": param.Id}).Updates(&order).RowsAffected
	return order.Id
}

// 获取订单列表
func (o *AppOrderService) GetList(param app.OrderQueryParam) []app.OrderList {
	//// 根据订单状态查询订单
	//orders := make([]app.Order, 0)
	//if param.Type == 1 {
	//	global.Db.Table("order").Where("status != ? and sid = ?", 5, param.Sid).Find(&orders)
	//} else {
	//	global.Db.Table("order").Where("status = ? and sid = ?", 5, param.Sid).Find(&orders)
	//}
	//
	//// 组装订单列表
	//orderList := make([]app.OrderList, 0)
	//for _, o := range orders {
	//	var order app.OrderList
	//	goods := make([]app.Goods, 0)
	//	order.Id = o.Id
	//	order.Status = o.Status
	//	order.TotalPrice = o.TotalPrice
	//	order.Created = o.Created
	//	// 查询服务信息
	//	goodsIds := make([]uint, 0)
	//	goodsIdsCount := map[int64]int{}
	//	for _, gidCount := range strings.Split(o.GoodsIdsCount, ",") {
	//		ic := strings.Split(gidCount, ":")
	//		if len(ic) < 2 {
	//			continue
	//		}
	//		id, _ := strconv.Atoi(ic[0])
	//		count, _ := strconv.Atoi(ic[1])
	//		goodsIdsCount[int64(id)] = count
	//		goodsIds = append(goodsIds, uint(id))
	//	}
	//	global.Db.Table("goods").Find(&goods, goodsIds)
	//
	//	// 组装服务项
	//	goodsItem := make([]app.GoodsItem, 0)
	//	for _, g := range goods {
	//		gItem := app.GoodsItem{
	//			Id:       g.Id,
	//			Title:    g.Title,
	//			Price:    g.Price,
	//			ImageUrl: g.ImageUrl,
	//			Count:    goodsIdsCount[int64(g.Id)],
	//		}
	//		order.GoodsCount = order.GoodsCount + uint(gItem.Count)
	//		goodsItem = append(goodsItem, gItem)
	//	}
	//	order.GoodsItem = goodsItem
	//	orderList = append(orderList, order)
	//}
	return nil
}

// 创建服务
func (o *AppOrderService) Create(c *gin.Context, param app.OrderCreateParam) uint64 {

	//查询服务信息
	consumerOpenId, _ := c.Get("openId")
	sid, _ := c.Get("sid")
	//查询发布者信息
	var servicesInfo app.Services
	global.Db.Table("services").Where(map[string]interface{}{"id": param.ServiceID}).First(&servicesInfo)
	if servicesInfo.Id <= 0 {
		return 0
	}
	servicesData := NewAppServicesService().GetInfo(servicesInfo.Id)
	serviceSnap, _ := json.Marshal(servicesData)
	orderNo := utils.GenerateOrderSn(servicesInfo.Id)

	order := app.Order{
		OrderNo:          orderNo,
		PublisherOpenId:  servicesInfo.OpenId,
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

func (o *AppOrderService) GetListByOpenId(param app.UserOrderQueryParam) ([]app.OrderInfo, int64) {
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
		rows = rowDao.Find(&[]app.Services{}).RowsAffected
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
		var servicesInfo app.ServicesInfo
		_ = json.Unmarshal([]byte(item.ServiceSnap), &servicesInfo)
		orderInfo := app.OrderInfo{
			ID:          item.Id,
			OrderNo:     item.OrderNo,
			Price:       servicesInfo.Price,
			ServiceSnap: servicesInfo,
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
	var servicesInfo app.ServicesInfo
	_ = json.Unmarshal([]byte(orderData.ServiceSnap), &servicesInfo)
	orderInfo := app.OrderInfo{
		ID:          orderData.Id,
		OrderNo:     orderData.OrderNo,
		Price:       servicesInfo.Price,
		ServiceSnap: servicesInfo,
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
func (o *AppOrderService) TodayData(param app.UserOrderTodayParam) app.OrderTodayDate {
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
