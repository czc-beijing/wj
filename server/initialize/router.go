package initialize

import (
	"fmt"
	"imall/api"
	"imall/global"
	"imall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	// 微信小程序API
	app := engine.Group("/homemaking")

	{
		//订单服务更新状态
		app.POST("v1/order/:id", api.GetAppOrder().UpdateOrderStatus)
		// 文件上传
		app.POST("/v1/file", api.GetWebFileUpload().FileUploadApp)
		// 开启JWT认证,以下接口需要认证成功才能访问
		app.GET("/v1/category", api.GetWebCategory().GetCategoryList)
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)
		// 用户登录
		app.GET("v1/token", api.GetAppUser().UserLogin)

		// 服务
		app.POST("/v1/service", api.GetAppServices().CreateServices)
		app.GET("/v1/service/list", api.GetAppServices().GetServicesList)
		app.GET("/v1/service/:id", api.GetAppServices().GetServicesInfo)
		app.POST("/v1/service/:id", api.GetAppServices().UpdateServicesStatus)
		app.PUT("/v1/service/:id", api.GetAppServices().UpdateServices)
		//评论
		app.GET("/v1/rating/service", api.GetAppRating().GetRatingList)

		app.Use(middleware.JWTAuth())
		//评论
		app.POST("/v1/rating", api.GetAppRating().CreateRating)
		app.GET("/v1/rating/order", api.GetAppRating().GetRatingInfo)

		// 用户登录
		app.POST("v1/user", api.GetAppUser().UserInfo)
		app.POST("v1/token/verify", api.GetAppUser().Verify)

		//我的预约服务
		app.GET("/v1/service/my", api.GetAppUser().GetUserServicesList)
		app.GET("/v1/service/count", api.GetAppUser().GetUserServicesTodayDate)
		//我的订单
		app.POST("v1/order", api.GetAppOrder().CreateOrder)
		app.GET("/v1/order/my", api.GetAppUser().GetUserOrderList)
		app.GET("/v1/order/:id", api.GetAppOrder().GetOrderDetail)
		app.GET("/v1/order/count", api.GetAppUser().GetUserOrderTodayDate)

	}
	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
