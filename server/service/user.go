package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"wj/common"
	"wj/global"
	"wj/models/app"
)

type AppUserService struct {
}

func (u *AppUserService) Login(context *gin.Context) (appUserInfo *app.UserInfo, err error) {
	code := context.Query("code")
	var acsJson app.Code2SessionResult
	sid := cast.ToInt64(context.Query("sid"))
	acs := app.Code2Session{
		Code:      code,
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res, err := http.DefaultClient.Get(fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code))
	if err != nil {
		return
	}
	err = json.NewDecoder(res.Body).Decode(&acsJson)
	if err != nil {
		return
	}
	if acsJson.ErrCode != 0 {
		err = errors.New("wx code error")
		return
	}
	rows := global.Db.Where(map[string]interface{}{"open_id": acsJson.OpenId, "sid": sid}).First(&app.User{}).RowsAffected
	if rows == 0 {
		user := app.User{
			OpenId:  acsJson.OpenId,
			Status:  1,
			Created: common.NowTime(),
			Sid:     sid,
		}
		row := global.Db.Create(&user).RowsAffected
		if row == 0 {
			return nil, errors.New("create Error")
		}
	}
	appUserInfo = &app.UserInfo{OpenId: acsJson.OpenId}
	return
}

func (u *AppUserService) SaveOrUpdate(c *gin.Context, userInfo app.UserInfoParam) *app.UserInfo {
	sid, _ := c.Get("sid")
	userInfo.Sid = cast.ToInt64(sid)
	openID, _ := c.Get("openId")
	userInfo.OpenId = openID.(string)
	// 查看用户是否已经存在
	user := &app.User{}
	rows := global.Db.Where(map[string]interface{}{"open_id": userInfo.OpenId, "sid": userInfo.Sid}).First(user).RowsAffected
	if rows == 0 {
		user = &app.User{
			Sid:      userInfo.Sid,
			OpenId:   userInfo.OpenId,
			Avatar:   userInfo.Avatar,
			Nickname: userInfo.Nickname,
			Status:   1,
			Created:  common.NowTime(),
		}
		_ = global.Db.Create(&user).RowsAffected
		return &app.UserInfo{
			Id:       user.Id,
			OpenId:   userInfo.OpenId,
			Avatar:   userInfo.Avatar,
			Nickname: userInfo.Nickname,
		}
	}
	if userInfo.Nickname == user.Nickname {
		return &app.UserInfo{
			Id:       user.Id,
			OpenId:   userInfo.OpenId,
			Avatar:   userInfo.Avatar,
			Nickname: userInfo.Nickname,
		}
	}
	updateUser := app.User{
		Sid:      userInfo.Sid,
		Avatar:   userInfo.Avatar,
		Nickname: userInfo.Nickname,
	}
	_ = global.Db.Model(&app.User{Id: user.Id}).Updates(&updateUser).RowsAffected
	return &app.UserInfo{
		Id:       user.Id,
		OpenId:   user.OpenId,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	}
}
