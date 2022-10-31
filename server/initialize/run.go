package initialize

import (
	"wj/log"
	"wj/utils"
)

func Run() {
	log.NewLog() //初始化日志
	LoadConfig()
	utils.SendSms()
	Mysql()
	Redis()
	go Cron()
	Router()
}
