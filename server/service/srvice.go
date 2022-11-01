package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
	"wj/common"
	"wj/constant"
	"wj/global"
	"wj/models"
	"wj/models/app"
	"wj/utils"
)

type WebServiceService struct {
}

type AppServiceService struct {
}

func NewAppServiceService() *AppServiceService {
	return &AppServiceService{}
}

// 获取服务列表
func (s *AppServiceService) GetList(c *gin.Context, param app.ServiceQueryParam) ([]app.ServiceInfo, int64) {
	sid, _ := c.Get("sid")
	param.Sid = cast.ToUint64(sid)

	var ServiceList []app.ServiceInfo
	var Service []app.Service
	rows := common.NewRestPage(param.Page, "service", &app.Service{
		CategoryId: param.CategoryId,
		Sid:        param.Sid,
		Status:     constant.ServiceStatusPublish,
	}, &Service, &[]app.Service{})
	if len(Service) <= 0 {
		return []app.ServiceInfo{}, 0
	}
	var openIDs []string
	var categoryIdIDs []uint64
	for _, good := range Service {
		openIDs = append(openIDs, good.OpenId)
		categoryIdIDs = append(categoryIdIDs, good.CategoryId)
	}
	openIDs = utils.RemoveDuplicationStr(openIDs)
	categoryIdIDs = utils.RemoveDuplicationUint64(categoryIdIDs)
	var categoryList []app.Category
	_ = global.Db.Table("category").Where("id IN ?", categoryIdIDs).Find(&categoryList).RowsAffected
	var userList []app.User
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	categoryMap := map[uint64]app.Category{}
	userMap := map[string]app.User{}
	for _, item := range categoryList {
		categoryMap[item.Id] = item
	}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	for _, good := range Service {
		ServiceInfo := app.ServiceInfo{
			ID:          good.Id,
			Type:        good.Type,
			Title:       good.Title,
			Description: good.Description,
			Price:       good.Price,
			Score:       good.Score,
			SalesVolume: good.SalesVolume,
			CreateTime:  good.CreateTime,
			BeginDate:   good.BeginDate,
			EndDate:     good.EndDate,
		}
		if good.DesignatedPlace > 0 {
			ServiceInfo.DesignatedPlace = true
		}
		if _, exists := categoryMap[good.CategoryId]; exists {
			ServiceInfo.Category.ID = categoryMap[good.CategoryId].Id
			ServiceInfo.Category.Name = categoryMap[good.CategoryId].Name
		}
		if _, exists := userMap[good.OpenId]; exists {
			ServiceInfo.Publisher.ID = userMap[good.OpenId].Id
			ServiceInfo.Publisher.Nickname = userMap[good.OpenId].Nickname
			ServiceInfo.Publisher.Avatar = userMap[good.OpenId].Avatar
			ServiceInfo.Publisher.Tel = good.Tel
		}
		ServiceInfo.CoverImage.ID = good.Id
		ServiceInfo.CoverImage.Path = good.CoverImage
		ServiceList = append(ServiceList, ServiceInfo)
	}
	return ServiceList, rows
}

// 获取服务详情
func (s *AppServiceService) GetInfo(id uint64) *app.ServiceInfo {
	var ServiceData *app.Service
	_ = global.Db.Table("service").Where("id = ?", id).First(&ServiceData).RowsAffected
	if ServiceData.Id == 0 {
		return nil
	}
	var categoryInfo *app.Category
	_ = global.Db.Table("category").Where("id = ?", ServiceData.CategoryId).First(&categoryInfo).RowsAffected
	var userInfo *app.User
	_ = global.Db.Table("user").Where("open_id = ?", ServiceData.OpenId).First(&userInfo).RowsAffected
	ServiceInfo := &app.ServiceInfo{
		ID:          ServiceData.Id,
		Type:        ServiceData.Type,
		Title:       ServiceData.Title,
		Description: ServiceData.Description,
		Price:       ServiceData.Price,
		Score:       ServiceData.Score,
		SalesVolume: ServiceData.SalesVolume,
		CreateTime:  ServiceData.CreateTime,
		Status:      ServiceData.Status,
		BeginDate:   ServiceData.BeginDate,
		EndDate:     ServiceData.EndDate,
	}
	if ServiceData.DesignatedPlace > 0 {
		ServiceInfo.DesignatedPlace = true
	}
	if categoryInfo != nil {
		ServiceInfo.Category.ID = categoryInfo.Id
		ServiceInfo.Category.Name = categoryInfo.Name
	}
	if userInfo != nil {
		ServiceInfo.Publisher.ID = userInfo.Id
		ServiceInfo.Publisher.Nickname = userInfo.Nickname
		ServiceInfo.Publisher.Avatar = userInfo.Avatar
		ServiceInfo.Publisher.Tel = ServiceData.Tel
	}
	ServiceInfo.CoverImage.ID = ServiceData.Id
	ServiceInfo.CoverImage.Path = ServiceData.CoverImage
	return ServiceInfo
}

// 创建服务
func (s *AppServiceService) Create(c *gin.Context, param app.ServiceCreateParam) uint64 {
	sid, _ := c.Get("sid")
	param.Sid = cast.ToUint64(sid)
	openID, _ := c.Get("openId")
	param.OpenId = openID.(string)
	designatedPlace := 0
	if param.DesignatedPlace {
		designatedPlace = 1
	}
	Service := app.Service{
		OpenId:          param.OpenId,
		CategoryId:      param.CategoryId,
		Type:            param.Type,
		Title:           param.Title,
		DesignatedPlace: designatedPlace,
		Description:     param.Description,
		Price:           param.Price,
		CoverImage:      param.CoverImage,
		Tel:             "13429208394",
		Sid:             param.Sid,
		BeginDate:       param.BeginDate,
		EndDate:         param.EndDate,
		CreateTime:      time.Now().Format("20060102"),
		UpdateTime:      time.Now().Format("20060102"),
		Status:          1,
		Score:           "100",
		SalesVolume:     10000,
	}
	_ = global.Db.Create(&Service).RowsAffected
	return Service.Id
}

// 更新服务状态
func (s *AppServiceService) UpdateStatus(param app.ServiceStatusUpdateParam) uint64 {
	Service := app.Service{
		Id:     param.Id,
		Status: param.Action,
	}
	_ = global.Db.Model(&Service).Update("status", Service.Status).RowsAffected
	return Service.Id
}

// 获取服务列表
func (s *AppServiceService) GetListByOpenId(c *gin.Context, appType int, page, pageSize int) ([]app.ServiceInfo, int64) {
	sid, _ := c.Get("sid")
	sid = cast.ToUint64(sid)
	openID, _ := c.Get("openId")
	openIDStr := openID.(string)
	status := c.Query("status")

	var ServiceList []app.ServiceInfo
	var Service []app.Service
	query := map[string]interface{}{"type": appType, "sid": sid, "open_id": openIDStr}
	if len(status) > 0 {
		query["status"] = cast.ToInt(status)
	}
	rows := common.NewRestPage(models.AppPage{
		PageNum:  page,
		PageSize: pageSize,
	}, "Service", query, &Service, &[]app.Service{})
	if len(Service) <= 0 {
		return []app.ServiceInfo{}, 0
	}
	var openIDs []string
	var categoryIdIDs []uint64
	openIDs = append(openIDs, openIDStr)
	for _, good := range Service {
		categoryIdIDs = append(categoryIdIDs, good.CategoryId)
	}
	openIDs = utils.RemoveDuplicationStr(openIDs)
	categoryIdIDs = utils.RemoveDuplicationUint64(categoryIdIDs)
	var categoryList []app.Category
	_ = global.Db.Table("category").Where("id IN ?", categoryIdIDs).Find(&categoryList).RowsAffected
	var userList []app.User
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	categoryMap := map[uint64]app.Category{}
	userMap := map[string]app.User{}
	for _, item := range categoryList {
		categoryMap[item.Id] = item
	}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	for _, good := range Service {
		ServiceInfo := app.ServiceInfo{
			ID:          good.Id,
			Type:        good.Type,
			Title:       good.Title,
			Description: good.Description,
			Price:       good.Price,
			Score:       good.Score,
			Status:      good.Status,
			SalesVolume: good.SalesVolume,
			CreateTime:  good.CreateTime,
			BeginDate:   good.BeginDate,
			EndDate:     good.EndDate,
		}
		if good.DesignatedPlace > 0 {
			ServiceInfo.DesignatedPlace = true
		}
		if _, exists := categoryMap[good.CategoryId]; exists {
			ServiceInfo.Category.ID = categoryMap[good.CategoryId].Id
			ServiceInfo.Category.Name = categoryMap[good.CategoryId].Name
		}
		if _, exists := userMap[good.OpenId]; exists {
			ServiceInfo.Publisher.ID = userMap[good.OpenId].Id
			ServiceInfo.Publisher.Nickname = userMap[good.OpenId].Nickname
			ServiceInfo.Publisher.Avatar = userMap[good.OpenId].Avatar
			ServiceInfo.Publisher.Tel = good.Tel
		}
		ServiceInfo.CoverImage.ID = good.Id
		ServiceInfo.CoverImage.Path = good.CoverImage
		ServiceList = append(ServiceList, ServiceInfo)
	}
	return ServiceList, rows
}

// 创建服务
func (s *AppServiceService) Update(c *gin.Context, param app.ServiceUpdateParam) uint64 {
	sid, _ := c.Get("sid")
	param.Sid = cast.ToUint64(sid)
	designatedPlace := 0
	if param.DesignatedPlace {
		designatedPlace = 1
	}
	updateService := app.Service{
		OpenId:          param.OpenId,
		CategoryId:      param.CategoryId,
		Type:            param.Type,
		Title:           param.Title,
		DesignatedPlace: designatedPlace,
		Description:     param.Description,
		Price:           param.Price,
		CoverImage:      param.CoverImage.Path,
		Tel:             "13429208394",
		Sid:             param.Sid,
		BeginDate:       param.BeginDate,
		EndDate:         param.EndDate,
		Score:           "100",
		SalesVolume:     10000,
		UpdateTime:      time.Now().Format("20060102"),
	}
	Service := app.Service{
		Id: param.Id,
	}
	_ = global.Db.Model(&Service).Updates(&updateService).RowsAffected
	return param.Id
}

// 统计当日数据
func (s *AppServiceService) TodayData(c *gin.Context, param app.UserServiceQueryParam) app.TodayDate {
	sid, _ := c.Get("sid")
	param.Sid = cast.ToUint64(sid)
	openID, _ := c.Get("openId")
	param.OpenId = openID.(string)
	pps := "open_id = ? and sid = ? and type = ? and status = 0"
	pds := "open_id = ? and sid = ? and type = ? and status = 1"
	prs := "open_id = ? and sid = ? and type = ? and status = 2"
	ras := "open_id = ? and sid = ? and type = ? and status = 3"
	ServiceDate := app.TodayDate{}
	global.Db.Table("service").Where(pps, param.OpenId, param.Sid, param.Type).Count(&ServiceDate.Pending)
	global.Db.Table("service").Where(pds, param.OpenId, param.Sid, param.Type).Count(&ServiceDate.Unpublished)
	global.Db.Table("service").Where(prs, param.OpenId, param.Sid, param.Type).Count(&ServiceDate.Published)
	global.Db.Table("service").Where(ras, param.OpenId, param.Sid, param.Type).Count(&ServiceDate.OffShelves)
	return ServiceDate
}
