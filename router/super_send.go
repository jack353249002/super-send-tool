package router

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/controller"
	"super-send-tool/middleware"
)

func SuperSendRouter(r *gin.Engine) {
	// 静态文件服务
	r.Static("/", "run_root")
	etcdBridge := r.Group("/etcd_bridge")
	{
		etcdBridge.Use(middleware.EtcdBridgeMiddInfo())
		etcdBridge.POST("/getEtcdBridgeList", controller.GetEtcdBridgeList)
		etcdBridge.POST("/addEtcdBridge", controller.AddEtcdBridge)
		etcdBridge.POST("/updateEtcdBridge", controller.UpdateEtcdBridge)
		etcdBridge.POST("/deleteEtcdBridge", controller.DeleteEtcdBridge)
		etcdBridge.POST("/loginEtcdBridge", controller.LoginEtcdBridge)
		etcdBridge.POST("/addUser", controller.AddEtcdBridgeUser)
		etcdBridge.POST("/updateUser", controller.UpdateEtcdBridgeUser)
		etcdBridge.POST("/delUser", controller.DelEtcdBridgeUser)
		etcdBridge.POST("/getUser", controller.GetEtcdBridgeUser)
		etcdBridge.POST("/getUserList", controller.GetEtcdBridgeUserList)
		etcdBridge.POST("/addEtcdBridgeRole", controller.AddEtcdBridgeRole)
		etcdBridge.POST("/getEtcdBridgeRolesList", controller.GetEtcdBridgeRolesList)
		etcdBridge.POST("/delEtcdBridgeRole", controller.DelEtcdBridgeRole)
		etcdBridge.POST("/addEtcdBridgeRoleForUser", controller.AddEtcdBridgeRoleForUser)
		etcdBridge.POST("/delEtcdBridgeRoleForUser", controller.DelEtcdBridgeRoleForUser)
		etcdBridge.POST("/getEtcdBridgeRolesPermissions", controller.GetEtcdBridgeRolesPermissions)
		etcdBridge.POST("/delEtcdBridgeRolesPermissions", controller.DelEtcdBridgeRolesPermissions)

		etcdBridge.POST("/addServer", controller.AddEtcdBridgeServer)
		etcdBridge.POST("/updateServer", controller.UpdateEtcdBridgeServer)
		etcdBridge.POST("/delServer", controller.DelEtcdBridgeServer)
		etcdBridge.POST("/getServer", controller.GetEtcdBridgeServer)
		etcdBridge.POST("/getServerList", controller.GetEtcdBridgeServerList)

		etcdBridge.POST("/addSyncConf", controller.AddEtcdBridgeSyncConf)
		etcdBridge.POST("/updateSyncConf", controller.UpdateEtcdBridgeSyncConf)
		etcdBridge.POST("/delSyncConf", controller.DelEtcdBridgeSyncConf)
		etcdBridge.POST("/getSyncConf", controller.GetEtcdBridgeSyncConf)
		etcdBridge.POST("/getSyncConfList", controller.GetEtcdBridgeSyncConfList)
		etcdBridge.POST("/setSyncAction", controller.EtcdBridgeSyncAction)
	}
	v1 := r.Group("/super_send")
	{
		v1.Use(middleware.MiddInfo())
		v1.POST("/loginTool", controller.LoginTool)
		v1.POST("/getSuperSendList", controller.SuperSendList)
		v1.POST("/getSuperSendInfo", controller.GetSuperSendInfo)
		v1.POST("/logoutUserAllDevice", controller.LogoutUserAllDevice)
		v1.POST("/addSuperSend", controller.AddSuperSend)
		v1.POST("/updateSuperSend", controller.UpdateSuperSend)
		v1.POST("/deleteSuperSend", controller.DeleteSuperSend)
		v1.POST("/loginSuperSend", controller.LoginSuperSend)
		v1.POST("/getOnlineSuperSend", controller.OnlineSuperSend)
		v1.POST("/setSuperSendOnline", controller.SetSuperSendOnline)
		v1.POST("/registerSuperSend", controller.RegisterSuperSend)
		v1.POST("/setMessage", controller.SetMessage)
		v1.POST("/delMessage", controller.DelMessage)
		v1.POST("/getMessageList", controller.GetMessageList)
		v1.POST("/setSend", controller.SetSend)
		v1.POST("/getSendInfoList", controller.GetSendInfoList)
		v1.POST("/getSendList", controller.GetSendList)
		v1.POST("/sendInfoAction", controller.SendInfoAction)
		v1.POST("/setSendInfo", controller.SetSendInfo)
		v1.POST("/getSmtpServerList", controller.GetSmtpServerList)
		v1.POST("/upload", controller.Upload)

		v1.POST("/getSmtpServer", controller.GetSmtpServer)
		v1.POST("/setSmtpServer", controller.SetSmtpServer)
		v1.POST("/delSmtpServer", controller.DelSmtpServer)
		v1.POST("/reloadServer", controller.ReloadServer)

		v1.POST("/setConf", controller.SetConf)
		v1.POST("/getConfList", controller.GetConfList)
		v1.POST("/delConf", controller.DelConf)
		v1.POST("/reloadConf", controller.ReloadConf)

		v1.POST("/addReceive", controller.AddReceive)
		v1.POST("/setReceive", controller.SetReceive)
		v1.POST("/getReceiveList", controller.GetReceiveList)
		v1.POST("/getReceiveMessagesList", controller.GetReceiveMessagesList)
		v1.POST("/getReceiveMessageInfo", controller.GetReceiveMessagesInfo)
		v1.POST("/receiveAction", controller.ReceiveAction)
		v1.POST("/addNewReceiveEmailMessageListen", controller.AddNewReceiveEmailMessageListen)

		v1.POST("/getImapServer", controller.GetImapServer)
		v1.POST("/setImapServer", controller.SetImapServer)
		v1.POST("/delImapServer", controller.DelImapServer)
		v1.POST("/getImapServerAllList", controller.GetImapServerAllList)
		v1.POST("/reloadImapServer", controller.ReloadImapServer)

		v1.POST("/createUser", controller.CreateUser)
		v1.POST("/getUsers", controller.GetUsers)
		v1.POST("/getUserInfo", controller.GetUserInfo)
		v1.POST("/delUser", controller.DelUser)
		v1.POST("/setPassword", controller.SetPassword)
		v1.POST("/setUserPassword", controller.SetUserPassword)

		v1.POST("/addRole", controller.AddRole)
		v1.POST("/getRolesList", controller.GetRolesList)
		v1.POST("/delRole", controller.DelRole)
		v1.POST("/addRoleForUser", controller.AddRoleForUser)
		v1.POST("/getRolesPermissions", controller.GetRolesPermissions)
		v1.POST("/delRolesPermissions", controller.DelRolesPermissions)
		v1.POST("/delUserRole", controller.DelRoleForUser)
	}
	checkAlive := r.Group("/check_alive")
	{
		checkAlive.Use(middleware.MiddInfo())
		checkAlive.POST("/create", controller.CreateCheckUserAlive)
		checkAlive.POST("/set", controller.SetCheckUserAlive)
		checkAlive.POST("/del", controller.DelCheckUserAlive)
		checkAlive.POST("/list", controller.CheckUserAliveList)
		checkAlive.POST("/ping", controller.CheckUserAlivePing)
	}
}
