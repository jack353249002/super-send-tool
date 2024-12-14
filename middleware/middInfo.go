package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"super-send-tool/controller"
)

var notAuth = map[string]struct{}{
	"/super_send/loginSuperSend":     {},
	"/super_send/getSuperSendList":   {},
	"/super_send/addSuperSend":       {},
	"/super_send/setSuperSendOnline": {},
	"/super_send/getOnlineSuperSend": {},
	"/super_send/updateSuperSend":    {},
	"/super_send/deleteSuperSend":    {},
}
var notSendID = map[string]struct{}{
	"/super_send/getSuperSendList":   {},
	"/super_send/addSuperSend":       {},
	"/super_send/setSuperSendOnline": {},
	"/super_send/getOnlineSuperSend": {},
	"/super_send/updateSuperSend":    {},
	"/super_send/deleteSuperSend":    {},
}

func MiddInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		superSendID := c.GetHeader("super_send_id")
		_, have := notSendID[c.Request.URL.Path]
		if superSendID == "" && !have {
			controller.ResponseFailed(c, nil, "super_send_id is required")
			c.Abort()
			return
		}
		_, have = notAuth[c.Request.URL.Path]
		if tokenString == "" && !have {
			controller.ResponseNotAuth(c, nil, "token is required")
			c.Abort()
			return
		}
		c.Set("token", tokenString)
		superSendIDInt, _ := strconv.Atoi(superSendID)
		c.Set("super_send_id", superSendIDInt)
		c.Next()
	}
}
