package main

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/config/baseconfig"
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
	r.Run(":8082") // 监听并在 0.0.0.0:8080 上启动服务
}
