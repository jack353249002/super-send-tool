package router

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/controller"
	"super-send-tool/middleware"
)

func SuperSendRouter(r *gin.Engine) {
	v1 := r.Group("/super_send")
	{
		v1.Use(middleware.MiddInfo())
		v1.POST("/getSuperSendList", controller.SuperSendList)
		v1.POST("/getSuperSendInfo", controller.GetSuperSendInfo)
		v1.POST("/addSuperSend", controller.AddSuperSend)
		v1.POST("/updateSuperSend", controller.UpdateSuperSend)
		v1.POST("/deleteSuperSend", controller.DeleteSuperSend)
		v1.POST("/loginSuperSend", controller.LoginSuperSend)
		v1.POST("/getOnlineSuperSend", controller.OnlineSuperSend)
		v1.POST("/setSuperSendOnline", controller.SetSuperSendOnline)
		v1.POST("/registerSuperSend", controller.RegisterSuperSend)
		v1.POST("/setMessage", controller.SetMessage)
		v1.POST("/getMessageList", controller.GetMessageList)
		v1.POST("/setSend", controller.SetSend)
		v1.POST("/getSendInfoList", controller.GetSendInfoList)
		v1.POST("/getSendList", controller.GetSendList)
		v1.POST("/sendInfoAction", controller.SendInfoAction)
		v1.POST("/setSendInfo", controller.SetSendInfo)
		v1.POST("/getSmtpServerList", controller.GetSmtpServerList)
		v1.POST("/upload", controller.Upload)
	}
}
