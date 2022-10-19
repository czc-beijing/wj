package api

import (
	"github.com/gin-gonic/gin"
	"wj/constant"
	"wj/models/app"
	"wj/response"
	"wj/service"
)

type AppRating struct {
	service.AppRatingService
}

func GetAppRating() *AppRating {
	return &AppRating{}
}

func (r *AppRating) GetRatingInfo(c *gin.Context) {
	var param app.RatingQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	ratingInfo := r.GetRating(c, param)
	response.Success(constant.Selected, ratingInfo, c)
}

func (r *AppRating) GetRatingList(c *gin.Context) {
	var param app.RatingListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if param.Page.PageSize == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	ratingList, rows := r.GetList(c, param)
	response.SuccessAppPage(constant.Selected, ratingList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (r *AppRating) CreateRating(c *gin.Context) {
	var param app.RatingCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if res := r.Create(c, param); res.Id > 0 {
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
