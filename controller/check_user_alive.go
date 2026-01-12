package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"super-send-tool/db"
	"super-send-tool/model/dbmodel"
	"super-send-tool/model/request"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 创建用户检测信息
// @Router /check_alive/create [POST]
func CreateCheckUserAlive(c *gin.Context) {
	var userRequest request.CreateUserAliveRequest
	err := c.ShouldBind(&userRequest)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewCheckUserAliveDao(c, db.Db)
	_, err = dao.Get(c, "*", "username=@username", map[string]interface{}{"username": userRequest.Username})
	if err != nil && err == gorm.ErrRecordNotFound {
		salt, errSalt := utils.GenerateSalt(8)
		if errSalt != nil {
			ResponseFailed(c, nil, "salt生成失败")
			return
		}
		userRequest.Password = utils.MD5(userRequest.Password, salt)
		err = dao.Create(c, &dbmodel.CheckUserAlive{Salt: salt, Username: userRequest.Username, Password: userRequest.Password, DayLoginFirstTime: userRequest.DayLoginFirstTime, SendID: userRequest.SendID, SendEmailActionTimeout: userRequest.SendEmailActionTimeout})
		if err != nil {
			ResponseFailed(c, nil, err.Error())
			return
		}
	} else if err != nil && err != gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseFailed(c, nil, "用户已存在")
		return
	}

	ResponseSuccess(c, nil, "success")
}

// @Summary 设置用户检测信息
// @Router /check_alive/set [POST]
func SetCheckUserAlive(c *gin.Context) {
	var userRequest request.SetUserAliveRequest
	err := c.ShouldBind(&userRequest)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewCheckUserAliveDao(c, db.Db)
	oldUser, err := dao.Get(c, "*", "id= @id", map[string]interface{}{"id": userRequest.ID})
	if err != nil && err == gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, "用户不存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, err.Error())
		return
	}

	_, err = dao.Get(c, "*", "username=@username AND id <> @id", map[string]interface{}{"username": userRequest.Username, "id": userRequest.ID})
	if err != nil && err == gorm.ErrRecordNotFound {
		saveData := make(map[string]interface{})
		saveData["username"] = userRequest.Username
		saveData["day_login_first_time"] = userRequest.DayLoginFirstTime
		saveData["send_id"] = userRequest.SendID
		saveData["send_email_action_timeout"] = userRequest.SendEmailActionTimeout
		if userRequest.Password != "" {
			saveData["password"] = utils.MD5(userRequest.Password, oldUser.Salt)
		}
		err = dao.Update(c, "id=@id", saveData, map[string]interface{}{"id": userRequest.ID})
		if err != nil {
			ResponseFailed(c, nil, err.Error())
			return
		}
	} else if err != nil && err != gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseFailed(c, nil, "用户已存在")
		return
	}

	ResponseSuccess(c, nil, "success")
}

// @Summary 删除用户
// @Router /check_alive/del [POST]
func DelCheckUserAlive(c *gin.Context) {
	var userRequest request.DelUserAliveRequest
	err := c.ShouldBind(&userRequest)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewCheckUserAliveDao(c, db.Db)
	err = dao.Delete(c, "id=@id", map[string]interface{}{"id": userRequest.ID})
	if err == nil {
		ResponseSuccess(c, nil, "success")
	} else {
		ResponseFailed(c, nil, err.Error())
	}
}

// @Summary 获取列表
// @Router /check_alive/list [POST]
func CheckUserAliveList(c *gin.Context) {
	dao := dbmodel.NewCheckUserAliveDao(c, db.Db)
	var pageListRequest request.CheckUserAliveRequest
	err := c.ShouldBind(&pageListRequest)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	offset := utils.CreatePageOffset(int(pageListRequest.PageNo), int(pageListRequest.PageSize))
	where := "1=@val"
	params := map[string]interface{}{"val": 1}
	if pageListRequest.KeyWords != "" {
		where += " AND ( username LIKE @keywords) "
		params["keywords"] = "%" + pageListRequest.KeyWords + "%"
	}
	list, err := dao.List(c, "*", where, offset, int(pageListRequest.PageSize), params)
	count, err := dao.Count(c, "id", where, params)
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		totalPage := utils.CreatePage(int(count), int(pageListRequest.PageSize))
		ResponseTableSuccess(c, list, int(pageListRequest.PageSize), int(pageListRequest.PageNo), totalPage, int(count), "success")
		return
	}
}

// @Summary 接收用户消息
// @Router /check_alive/ping [POST]
func CheckUserAlivePing(c *gin.Context) {
	dao := dbmodel.NewCheckUserAliveDao(c, db.Db)
	var request request.UserAlivePing
	err := c.ShouldBind(&request)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	user, err := dao.Get(c, "*", "username=@username", map[string]interface{}{"username": request.Username})
	if err != nil && err == gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, "用户不存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		password := utils.MD5(request.Password, user.Salt)
		if password != user.Password {
			ResponseFailed(c, nil, "密码错误")
			return
		} else {
			nowTime := time.Now()
			nowDayBegin := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
			oldUserTime := time.Unix(int64(user.DayLoginFirstTime), 0)
			oldUserTimeBegin := time.Date(oldUserTime.Year(), oldUserTime.Month(), oldUserTime.Day(), 0, 0, 0, 0, oldUserTime.Location())
			if nowDayBegin.Unix() == oldUserTimeBegin.Unix() {
				ResponseSuccess(c, nil, "success")
				return
			}
			dao.Update(c, "id=@id", map[string]interface{}{"day_login_first_time": nowTime.Unix()}, map[string]interface{}{"id": user.ID})
			ResponseSuccess(c, nil, "success")
		}
	}
}
