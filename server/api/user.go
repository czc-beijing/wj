package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"imall/common"
	"imall/constant"
	"imall/middleware"
	"imall/models/app"
	"imall/models/web"
	"imall/response"
	"imall/service"
	"net/http"
)

type WebUser struct {
	service.UserService
}

type AppUser struct {
	service.AppUserService
}

func GetWebUser() *WebUser {
	return &WebUser{}
}

func GetAppUser() *AppUser {
	return &AppUser{}
}

// 获取验证码
func (u *WebUser) GetCaptcha(c *gin.Context) {
	id, base64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": base64s}
	response.Success("操作成功", data, c)
}

// 用户登录
func (u *WebUser) UserLogin(c *gin.Context) {
	var param web.LoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := u.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		userInfo := web.UserInfo{
			Sid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

func (u *AppUser) UserLogin(context *gin.Context) {
	code := context.Query("code")
	if code == "" {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	userInfo := u.Login(code)
	if userInfo == nil {
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
	openID, exists := c.Get("openId")
	if !exists {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	sid, exists := c.Get("sid")
	if exists {
		param.Sid = cast.ToInt64(sid)
	}
	param.OpenId = openID.(string)
	if userInfo := u.UpdateUserInfo(param); userInfo != nil {
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

func (u *AppUser) GetUserServicesList(c *gin.Context) {
	var param app.UserServicesQueryParam
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
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	status := c.Query("status")
	servicesList, rows := service.NewAppServicesService().GetListByOpenId(param.Type, status, param.Page.PageNum, param.Page.PageSize, param.Sid, param.OpenId)
	response.SuccessAppPage(constant.Selected, servicesList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (u *AppUser) GetUserServicesTodayDate(c *gin.Context) {
	var param app.UserServicesQueryParam
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
	data := service.NewAppServicesService().TodayData(param)
	response.Success(constant.LoginSuccess, data, c)
}

func (u *AppUser) GetUserOrderList(c *gin.Context) {
	var param app.UserOrderQueryParam
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
	if param.Page.PageNum == 0 {
		param.Page.PageSize = 0
		param.Page.PageNum = 4
	}
	orderList, rows := service.NewAppOrderService().GetListByOpenId(param)
	response.SuccessAppPage(constant.Selected, orderList, rows, param.Page.PageNum, param.Page.PageSize, c)
}

func (u *AppUser) GetUserOrderTodayDate(c *gin.Context) {
	var param app.UserOrderTodayParam
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
	data := service.NewAppOrderService().TodayData(param)
	response.Success(constant.LoginSuccess, data, c)
}
