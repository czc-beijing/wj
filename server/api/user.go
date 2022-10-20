package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wj/constant"
	"wj/middleware"
	"wj/models/app"
	"wj/response"
	"wj/service"
)

type AppUser struct {
	service.AppUserService
}

func GetAppUser() *AppUser {
	return &AppUser{}
}

func (u *AppUser) UserLogin(context *gin.Context) {
	if len(context.Query("code")) <= 0 {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	sid, exists := context.Get("sid")
	if !exists || sid == 0 {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	userInfo, err := u.Login(context)
	if userInfo == nil || err != nil {
		response.AppFailed(http.StatusUnauthorized, 401, constant.StatusUnauthorized, context)
		return
	}
	token, _ := middleware.NewJWT().CreateToken(userInfo.OpenId)
	userInfo.Token = token
	response.Success(constant.LoginSuccess, userInfo, context)
	return
}

func (u *AppUser) UserInfo(c *gin.Context) {
	var param app.UserInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	_, exists := c.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if !exists || sid == 0 {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if userInfo := u.SaveOrUpdate(c, param); userInfo != nil {
		response.Success(constant.LoginSuccess, userInfo, c)
		return
	}
	response.Success(constant.LoginSuccess, nil, c)
	return
}

func (u *AppUser) Verify(c *gin.Context) {
	data := gin.H{
		"valid": true,
	}
	response.Success(constant.LoginSuccess, data, c)
	return
}
