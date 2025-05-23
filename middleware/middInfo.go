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
	"/super_send/getSuperSendInfo":   {},
}
var notSendID = map[string]struct{}{
	"/super_send/getSuperSendList":   {},
	"/super_send/addSuperSend":       {},
	"/super_send/setSuperSendOnline": {},
	"/super_send/getOnlineSuperSend": {},
	"/super_send/updateSuperSend":    {},
	"/super_send/deleteSuperSend":    {},
	"/super_send/getSuperSendInfo":   {},
}
var notEtcdBridgeID = map[string]struct{}{}

func MiddInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		superSendID := c.GetHeader("Super_send_id")
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
func EtcdBridgeMiddInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		etcdBridgeID := c.GetHeader("Etcd_bridge_id")
		_, have := notEtcdBridgeID[c.Request.URL.Path]
		if etcdBridgeID == "" && !have {
			controller.ResponseFailed(c, nil, "etcd_bridge_id is required")
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
		etcdBridgeIDInt, _ := strconv.Atoi(etcdBridgeID)
		c.Set("etcd_bridge_id", etcdBridgeIDInt)
		c.Next()
	}
}
