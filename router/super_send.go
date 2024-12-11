package router

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/controller"
)

func SuperSendRouter(r *gin.Engine) {
	v1 := r.Group("/super_send")
	{
		v1.POST("/getSuperSendList", controller.SuperSendList)
		v1.POST("/addSuperSend", controller.AddSuperSend)
		v1.POST("/updateSuperSend", controller.UpdateSuperSend)
		v1.POST("/deleteSuperSend", controller.DeleteSuperSend)
		v1.POST("/loginSuperSend", controller.LoginSuperSend)
		v1.POST("/getOnlineSuperSend", controller.OnlineSuperSend)
		v1.POST("/setSuperSendOnline", controller.SetSuperSendOnline)
	}
}
