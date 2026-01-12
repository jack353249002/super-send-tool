package main

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/config/baseconfig"
	"super-send-tool/cron"
	"super-send-tool/db"
	"super-send-tool/install"
	"super-send-tool/middleware"
	"super-send-tool/router"
)

func main() {
	db.SqLiteSetup(baseconfig.CONFIG.DBPath)
	install.InstallDataBase()
	r := gin.Default()
	middleware.Cors(r)
	router.SuperSendRouter(r)
	go cron.SendEmailCron()
	if baseconfig.CONFIG.OpenSSL {
		// 启动 HTTPS 服务
		err := r.RunTLS(baseconfig.CONFIG.Host, baseconfig.CONFIG.ServerPublicKey, baseconfig.CONFIG.ServerPrivateKey)
		if err != nil {
			panic(err)
		}
	} else {
		r.Run(baseconfig.CONFIG.Host) // 监听并在 0.0.0.0:8080 上启动服务
	}
}
