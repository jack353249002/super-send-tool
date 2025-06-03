package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"super-send-tool/controller"
	"super-send-tool/db"
	"super-send-tool/grpccons"
	"super-send-tool/model/dbmodel"
	"super-send-tool/proto"
	"time"
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
var notEtcdBridgeAUth = map[string]struct{}{
	"/etcd_bridge/getEtcdBridgeList": {},
}
var notEtcdBridgeID = map[string]struct{}{
	"/etcd_bridge/getEtcdBridgeList": {},
}

func MiddInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ""
		//tokenString := c.GetHeader("Token")
		superSendID := c.GetHeader("Super_send_id")
		superSendIDInt, _ := strconv.Atoi(superSendID)
		_, haveNotSendID := notSendID[c.Request.URL.Path]
		if superSendID == "" && !haveNotSendID {
			controller.ResponseFailed(c, nil, "super_send_id is required")
			c.Abort()
			return
		}
		_, haveNotAuth := notAuth[c.Request.URL.Path]
		if !haveNotSendID && !haveNotAuth {
			dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
			superSendInfo, err := dao.Get(c, "*", "id=@id", map[string]interface{}{"id": superSendIDInt})
			tokenString = superSendInfo.Token
			if err != nil {
				controller.ResponseFailed(c, nil, err.Error())
				c.Abort()
				return
			} else if superSendInfo.RefreshTokenInterval != 0 {
				nowTime := time.Now().Unix()
				if nowTime-int64(superSendInfo.ConnLastLoginTime) > int64(superSendInfo.RefreshTokenInterval) {
					con := grpccons.SuperSendGroupAction.Get(superSendIDInt)
					if con == nil {
						controller.ResponseFailed(c, nil, "连接失败")
						return
					}
					client := proto.NewUsersServiceClient(con.Conn)
					res, err := client.Login(c, &proto.LoginRequest{Username: superSendInfo.Username, Password: superSendInfo.Password})
					if err != nil {
						controller.ResponseFailed(c, nil, err.Error())
						return
					} else {
						if res.Code == 1 {
							dao.Update(c, "id=@id", map[string]interface{}{"token": res.GetLoginResponse().GetToken(), "conn_last_login_time": nowTime}, map[string]interface{}{"id": superSendIDInt})
							tokenString = res.GetLoginResponse().GetToken()
						} else {
							controller.ResponseFailed(c, nil, res.Message)
						}
					}
				}
			}
		}
		c.Set("token", tokenString)
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
		_, have = notEtcdBridgeAUth[c.Request.URL.Path]
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
