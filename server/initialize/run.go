package initialize

import "imall/log"

func Run() {
	log.NewLog() //初始化日志
	LoadConfig()
	Mysql()
	Redis()
	go Cron()
	Router()
}
