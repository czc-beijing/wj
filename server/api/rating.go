package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"imall/constant"
	"imall/models/app"
	"imall/response"
	"imall/service"
)

type WebRating struct {
	service.WebRatingService
}

type AppRating struct {
	service.AppRatingService
}

func GetWebRating() *WebRating {
	return &WebRating{}
}

func GetAppRating() *AppRating {
	return &AppRating{}
}

func (g *AppRating) GetRatingInfo(c *gin.Context) {
	var param app.RatingQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	openID, exists := c.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	param.OpenId = openID.(string)
	ratingInfo := service.NewAppRatingService().GetRatingInfo(param)
	response.Success(constant.Selected, ratingInfo, c)
}

func (g *AppRating) GetRatingList(c *gin.Context) {
	var param app.RatingListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	if param.Page.PageSize == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	ratingList, rows := g.GetList(param)
	response.SuccessAppPage(constant.Selected, ratingList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (g *AppRating) CreateRating(c *gin.Context) {
	var param app.RatingCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToUint64(sid)
	}
	openID, exists := c.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	param.OpenId = openID.(string)
	if res := service.NewAppRatingService().CreateRating(param); res.Id > 0 {
		//更新订单状态
		_ = service.NewAppOrderService().UpdateStatus(app.OrderUpdateParam{
			Id:     param.OrderId,
			Action: 5,
		})
		response.Success(constant.Created, res, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}
