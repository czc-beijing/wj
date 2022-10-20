package initialize

import (
	"fmt"
	"net/http"
	"wj/api"
	"wj/global"
	"wj/middleware"

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
		app.GET("/v1/category", api.GetAppCategory().GetCategoryList)
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)
		// 用户登录
		app.GET("v1/token", api.GetAppUser().UserLogin)
		// 文件上传
		app.POST("/v1/file", api.GetWebFileUpload().FileUploadApp)
		app.GET("/v1/service/list", api.GetAppService().GetServiceList)
		app.GET("/v1/service/:id", api.GetAppService().GetServiceInfo)

		// 服务
		app.Use(middleware.JWTAuth())
		app.POST("/v1/service", api.GetAppService().CreateService)
		app.POST("/v1/service/:id", api.GetAppService().UpdateServiceStatus)
		app.PUT("/v1/service/:id", api.GetAppService().UpdateService)
		//评论
		app.POST("/v1/rating", api.GetAppRating().CreateRating)
		app.GET("/v1/rating/order", api.GetAppRating().GetRatingInfo)
		app.GET("/v1/rating/service", api.GetAppRating().GetRatingList)

		// 用户登录
		app.POST("v1/user", api.GetAppUser().UserInfo)
		app.POST("v1/token/verify", api.GetAppUser().Verify)

		//我的预约服务
		app.GET("/v1/service/my", api.GetAppService().MyServiceList)
		app.GET("/v1/service/count", api.GetAppService().MyServiceTodayDate)

		//我的订单
		//订单服务更新状态
		app.POST("v1/order/:id", api.GetAppOrder().UpdateOrderStatus)
		app.POST("v1/order", api.GetAppOrder().CreateOrder)
		app.GET("/v1/order/my", api.GetAppOrder().GetUserOrderList)
		app.GET("/v1/order/:id", api.GetAppOrder().GetOrderDetail)
		app.GET("/v1/order/count", api.GetAppOrder().GetUserOrderTodayDate)

	}
	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
