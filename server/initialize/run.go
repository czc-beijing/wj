package initialize

import "wj/log"

func Run() {
	log.NewLog() //初始化日志
	LoadConfig()
	Mysql()
	Redis()
	go Cron()
	Router()
}
