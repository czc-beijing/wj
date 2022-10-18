package service

import (
	"github.com/spf13/cast"
	"imall/common"
	"imall/constant"
	"imall/global"
	"imall/models"
	"imall/models/app"
	"imall/utils"
	"time"
)

type WebServicesService struct {
}

type AppServicesService struct {
}

func NewAppServicesService() *AppServicesService {
	return &AppServicesService{}
}

// 获取服务列表
func (s *AppServicesService) GetList(param app.ServicesQueryParam) ([]app.ServicesInfo, int64) {
	var servicesList []app.ServicesInfo
	var services []app.Services
	rows := common.NewRestPage(param.Page, "services", &app.Services{
		CategoryId: param.CategoryId,
		Sid:        param.Sid,
		Status:     constant.ServicesStatusPublish,
	}, &services, &[]app.Services{})
	if len(services) <= 0 {
		return []app.ServicesInfo{}, 0
	}
	var openIDs []string
	var categoryIdIDs []uint64
	for _, good := range services {
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
	for _, good := range services {
		servicesInfo := app.ServicesInfo{
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
			servicesInfo.DesignatedPlace = true
		}
		if _, exists := categoryMap[good.CategoryId]; exists {
			servicesInfo.Category.ID = categoryMap[good.CategoryId].Id
			servicesInfo.Category.Name = categoryMap[good.CategoryId].Name
		}
		if _, exists := userMap[good.OpenId]; exists {
			servicesInfo.Publisher.ID = userMap[good.OpenId].Id
			servicesInfo.Publisher.Nickname = userMap[good.OpenId].Nickname
			servicesInfo.Publisher.Avatar = userMap[good.OpenId].Avatar
			servicesInfo.Publisher.Tel = good.Tel
		}
		servicesInfo.CoverImage.ID = good.Id
		servicesInfo.CoverImage.Path = good.CoverImage
		servicesList = append(servicesList, servicesInfo)
	}
	return servicesList, rows
}

// 获取服务详情
func (s *AppServicesService) GetInfo(id uint64) *app.ServicesInfo {
	var servicesData *app.Services
	_ = global.Db.Table("services").Where("id = ?", id).First(&servicesData).RowsAffected
	if servicesData.Id == 0 {
		return nil
	}
	var categoryInfo *app.Category
	_ = global.Db.Table("category").Where("id = ?", servicesData.CategoryId).First(&categoryInfo).RowsAffected
	var userInfo *app.User
	_ = global.Db.Table("user").Where("open_id = ?", servicesData.OpenId).First(&userInfo).RowsAffected
	servicesInfo := &app.ServicesInfo{
		ID:          servicesData.Id,
		Type:        servicesData.Type,
		Title:       servicesData.Title,
		Description: servicesData.Description,
		Price:       servicesData.Price,
		Score:       servicesData.Score,
		SalesVolume: servicesData.SalesVolume,
		CreateTime:  servicesData.CreateTime,
		Status:      servicesData.Status,
		BeginDate:   servicesData.BeginDate,
		EndDate:     servicesData.EndDate,
	}
	if servicesData.DesignatedPlace > 0 {
		servicesInfo.DesignatedPlace = true
	}
	if categoryInfo != nil {
		servicesInfo.Category.ID = categoryInfo.Id
		servicesInfo.Category.Name = categoryInfo.Name
	}
	if userInfo != nil {
		servicesInfo.Publisher.ID = userInfo.Id
		servicesInfo.Publisher.Nickname = userInfo.Nickname
		servicesInfo.Publisher.Avatar = userInfo.Avatar
		servicesInfo.Publisher.Tel = servicesData.Tel
	}
	servicesInfo.CoverImage.ID = servicesData.Id
	servicesInfo.CoverImage.Path = servicesData.CoverImage
	return servicesInfo
}

// 创建服务
func (s *AppServicesService) Create(param app.ServicesCreateParam) uint64 {
	designatedPlace := 0
	if param.DesignatedPlace {
		designatedPlace = 1
	}
	Services := app.Services{
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
	_ = global.Db.Create(&Services).RowsAffected
	return Services.Id
}

// 更新服务状态
func (s *AppServicesService) UpdateStatus(param app.ServicesStatusUpdateParam) uint64 {
	Services := app.Services{
		Id:     param.Id,
		Status: param.Action,
	}
	_ = global.Db.Model(&Services).Update("status", Services.Status).RowsAffected
	return Services.Id
}

// 获取服务列表
func (s *AppServicesService) GetListByOpenId(appType int, status string, page, pageSize int, sid uint64, openId string) ([]app.ServicesInfo, int64) {
	var servicesList []app.ServicesInfo
	var services []app.Services
	query := map[string]interface{}{"type": appType, "sid": sid, "open_id": openId}
	if len(status) > 0 {
		query["status"] = cast.ToInt(status)
	}
	rows := common.NewRestPage(models.AppPage{
		PageNum:  page,
		PageSize: pageSize,
	}, "services", query, &services, &[]app.Services{})
	if len(services) <= 0 {
		return []app.ServicesInfo{}, 0
	}
	var openIDs []string
	var categoryIdIDs []uint64
	openIDs = append(openIDs, openId)
	for _, good := range services {
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
	for _, good := range services {
		servicesInfo := app.ServicesInfo{
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
			servicesInfo.DesignatedPlace = true
		}
		if _, exists := categoryMap[good.CategoryId]; exists {
			servicesInfo.Category.ID = categoryMap[good.CategoryId].Id
			servicesInfo.Category.Name = categoryMap[good.CategoryId].Name
		}
		if _, exists := userMap[good.OpenId]; exists {
			servicesInfo.Publisher.ID = userMap[good.OpenId].Id
			servicesInfo.Publisher.Nickname = userMap[good.OpenId].Nickname
			servicesInfo.Publisher.Avatar = userMap[good.OpenId].Avatar
			servicesInfo.Publisher.Tel = good.Tel
		}
		servicesInfo.CoverImage.ID = good.Id
		servicesInfo.CoverImage.Path = good.CoverImage
		servicesList = append(servicesList, servicesInfo)
	}
	return servicesList, rows
}

// 创建服务
func (s *AppServicesService) Update(param app.ServicesUpdateParam) uint64 {
	designatedPlace := 0
	if param.DesignatedPlace {
		designatedPlace = 1
	}
	updateServices := app.Services{
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
	services := app.Services{
		Id: param.Id,
	}
	_ = global.Db.Model(&services).Updates(&updateServices).RowsAffected
	return param.Id
}

// 统计当日数据
func (s *AppServicesService) TodayData(param app.UserServicesQueryParam) app.TodayDate {
	pps := "open_id = ? and sid = ? and type = ? and status = 0"
	pds := "open_id = ? and sid = ? and type = ? and status = 1"
	prs := "open_id = ? and sid = ? and type = ? and status = 2"
	ras := "open_id = ? and sid = ? and type = ? and status = 3"
	serviceDate := app.TodayDate{}
	global.Db.Table("services").Where(pps, param.OpenId, param.Sid, param.Type).Count(&serviceDate.Pending)
	global.Db.Table("services").Where(pds, param.OpenId, param.Sid, param.Type).Count(&serviceDate.Unpublished)
	global.Db.Table("services").Where(prs, param.OpenId, param.Sid, param.Type).Count(&serviceDate.Published)
	global.Db.Table("services").Where(ras, param.OpenId, param.Sid, param.Type).Count(&serviceDate.OffShelves)
	return serviceDate
}
