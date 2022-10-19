package constant

const (
	ServiceStatusInit       int = 0
	ServiceStatusNotPublish int = 1 //待发布
	ServiceStatusPublish    int = 2 //发布
	ServiceStatusOffline    int = 3 //下线
	ServiceStatusCancel     int = 4 //取消
	ServiceStatusNotPass    int = 5 //不通过
)
