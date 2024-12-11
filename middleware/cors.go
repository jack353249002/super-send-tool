package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors(r *gin.Engine) {
	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                       // 允许的来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的 HTTP 方法
	config.AllowHeaders = []string{"*"}                                       // 允许的请求头
	config.ExposeHeaders = []string{"X-Total-Count"}                          // 允许暴露的响应头
	config.AllowCredentials = false                                           // 是否允许发送 Cookie
	config.MaxAge = 12 * time.Hour                                            // 预检请求的有效期

	r.Use(cors.New(config))
}
