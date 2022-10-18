package service

import (
	"encoding/json"
	"github.com/spf13/cast"
	"imall/common"
	"imall/global"
	"imall/models/app"
)

type WebRatingService struct {
}

type AppRatingService struct {
}

func NewAppRatingService() *AppRatingService {
	return &AppRatingService{}
}

// 获取服务列表
func (r *AppRatingService) GetList(param app.RatingListParam) ([]app.RatingInfo, int64) {
	var ratingData []*app.Rating
	query := map[string]interface{}{"sid": param.Sid, "service_id": param.ServiceId}
	rows := common.NewRestPage(param.Page, "rating", &app.Rating{
		ServiceId: param.ServiceId,
		Sid:       param.Sid,
	}, &ratingData, &[]app.Services{})
	if rows <= 0 {
		return []app.RatingInfo{}, 0
	}
	_ = global.Db.Table("rating").Where(query).Find(&ratingData).RowsAffected
	var userList []app.User
	var openIDs []string
	for _, item := range ratingData {
		openIDs = append(openIDs, item.OpenId)
	}
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	userMap := map[string]app.User{}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	var ratingList []app.RatingInfo
	for _, item := range ratingData {
		var arrIllustration []string
		ratingInfo := app.RatingInfo{
			ID:         item.Id,
			Score:      item.Score,
			Content:    item.Content,
			CreateTime: item.Created,
		}
		_ = json.Unmarshal([]byte(item.Illustration), &arrIllustration)
		ratingInfo.Illustration = arrIllustration
		if userInfo, exists := userMap[item.OpenId]; exists {
			ratingInfo.Author.ID = userInfo.Id
			ratingInfo.Author.Nickname = userInfo.Nickname
			ratingInfo.Author.Avatar = userInfo.Avatar
		}
		ratingList = append(ratingList, ratingInfo)
	}
	return ratingList, rows
}

func (r *AppRatingService) CreateRating(param app.RatingCreateParam) app.Rating {
	var orderInfo app.Order
	_ = global.Db.Table("order").Where(map[string]interface{}{"id": param.OrderId}).First(&orderInfo).RowsAffected
	Illustration, _ := json.Marshal(param.Illustration)
	rating := app.Rating{
		OrderId:      param.OrderId,
		ServiceId:    orderInfo.ServiceID,
		Score:        param.Score,
		Content:      param.Content,
		OpenId:       param.OpenId,
		Illustration: cast.ToString(Illustration),
		Sid:          param.Sid,
	}
	_ = global.Db.Create(&rating).RowsAffected
	return rating
}

// 获取服务列表
func (r *AppRatingService) GetRatingInfo(param app.RatingQueryParam) app.RatingInfo {
	var ratingData *app.Rating
	query := map[string]interface{}{"open_id": param.OpenId, "sid": param.Sid, "order_id": param.OrderId}
	_ = global.Db.Table("rating").Where(query).First(&ratingData).RowsAffected
	var userList []app.User
	var openIDs []string
	openIDs = append(openIDs, param.OpenId)
	_ = global.Db.Table("user").Where("open_id IN ?", openIDs).Find(&userList).RowsAffected
	userMap := map[string]app.User{}
	for _, item := range userList {
		userMap[item.OpenId] = item
	}
	var arrIllustration []string
	ratingInfo := app.RatingInfo{
		ID:         ratingData.Id,
		Score:      ratingData.Score,
		Content:    ratingData.Content,
		CreateTime: ratingData.Created,
	}
	_ = json.Unmarshal([]byte(ratingData.Illustration), &arrIllustration)
	ratingInfo.Illustration = arrIllustration
	if userInfo, exists := userMap[ratingData.OpenId]; exists {
		ratingInfo.Author.ID = userInfo.Id
		ratingInfo.Author.Nickname = userInfo.Nickname
		ratingInfo.Author.Avatar = userInfo.Avatar
	}
	return ratingInfo
}
