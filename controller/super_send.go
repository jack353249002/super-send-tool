package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
	"super-send-tool/db"
	"super-send-tool/grpccons"
	"super-send-tool/model/dbmodel"
	"super-send-tool/model/request"
	"super-send-tool/model/response"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 连接和断开
// @Router /super_send/setSuperSendOnline [POST]
func SetSuperSendOnline(c *gin.Context) {
	var setSuperSendOnlineRequest request.SetSuperSendOnlineRequest
	err := c.ShouldBind(&setSuperSendOnlineRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	saveData := map[string]interface{}{"online": setSuperSendOnlineRequest.Online}
	err = dao.Update(c, "id=@id", saveData, map[string]interface{}{"id": setSuperSendOnlineRequest.ID})
	if err != nil {
		if setSuperSendOnlineRequest.Online == 0 {
			con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
			if con == nil {
				ResponseFailed(c, nil, "连接失败")
				return
			}
			client := proto.NewUsersServiceClient(con.Conn)
			// 创建元数据
			md := metadata.Pairs(
				"token", c.GetString("token"),
			)
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			ctx = metadata.NewOutgoingContext(ctx, md)
			client.LogOut(ctx, &empty.Empty{})
		}
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 断开所有设备
// @Router /super_send/logoutUserAllDevice [POST]
func LogoutUserAllDevice(c *gin.Context) {
	var logOutSendOnlineRequest request.LogOutUserAllDeviceRequest
	err := c.ShouldBind(&logOutSendOnlineRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	client.LogOutAll(ctx, &empty.Empty{})
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 获取在线的列表
// @Router /super_send/getOnlineSuperSend [POST]
func OnlineSuperSend(c *gin.Context) {
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	list, err := dao.ListAll(c, "*", "online=@val", map[string]interface{}{"val": 1})
	if err == nil {
		var res []response.OnlineMenu
		for _, val := range list {
			var row response.OnlineMenu
			row.ID = val.ID
			row.Title = fmt.Sprintf("id:%d,address:%s,username:%s", val.ID, val.Address, val.Username)
			row.Name = fmt.Sprintf("%d-%s", val.ID, val.Address)
			row.Token = val.Token
			res = append(res, row)
		}
		ResponseSuccess(c, res, "success")
		return
	} else {
		ResponseFailed(c, nil, err.Error())
		return
	}
}

// @Summary 获取列表
// @Router /super_send/getSuperSendList [POST]
func SuperSendList(c *gin.Context) {
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	var pageListRequest request.SuperSendListRequest
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
		where += " AND (address LIKE @keywords OR username LIKE @keywords) "
		params["keywords"] = "%" + pageListRequest.KeyWords + "%"
	}
	if pageListRequest.IsSSL > -1 {
		where += " AND is_ssl=@is_ssl "
		params["is_ssl"] = pageListRequest.IsSSL
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

// @Summary 获取信息
// @Router /super_send/getSuperSendInfo [POST]
func GetSuperSendInfo(c *gin.Context) {
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	var getSuperSendInfo request.GetSuperSendInfoRequest
	err := c.ShouldBind(&getSuperSendInfo)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	info, err := dao.Get(c, "*", "id=@id", map[string]interface{}{"id": getSuperSendInfo.ID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseSuccess(c, info, "success")
		return
	}
}

// @Summary 添加super_send
// @Router /super_send/addSuperSend [POST]
func AddSuperSend(c *gin.Context) {
	var addSuperSendRequest request.AddSuperSendRequest
	err := c.ShouldBind(&addSuperSendRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	err = dao.Create(c, &dbmodel.SuperSendConnInfo{Address: addSuperSendRequest.Address, IsSsl: int(addSuperSendRequest.IsSSL), Password: addSuperSendRequest.Password, Username: addSuperSendRequest.UserName})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 编辑super_send
// @Router /super_send/updateSuperSend [POST]
func UpdateSuperSend(c *gin.Context) {
	var updateSuperSendRequest request.UpdateSuperSendRequest
	err := c.ShouldBind(&updateSuperSendRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	saveData := map[string]interface{}{"address": updateSuperSendRequest.Address, "is_ssl": updateSuperSendRequest.IsSSL, "password": updateSuperSendRequest.Password, "username": updateSuperSendRequest.UserName}
	err = dao.Update(c, "id=@id", saveData, map[string]interface{}{"id": updateSuperSendRequest.ID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 删除super_send
// @Router /super_send/deleteSuperSend [POST]
func DeleteSuperSend(c *gin.Context) {
	var deleteSuperSendRequest request.DeleteSuperSendRequest
	err := c.ShouldBind(&deleteSuperSendRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	err = dao.Delete(c, "id=@id", map[string]interface{}{"id": deleteSuperSendRequest.ID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 登录super_send
// @Router /super_send/loginSuperSend [POST]
func LoginSuperSend(c *gin.Context) {
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	superSendInfo, err := dao.Get(c, "*", "id=@id", map[string]interface{}{"id": c.GetInt("super_send_id")})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
		if con == nil {
			ResponseFailed(c, nil, "连接失败")
			return
		}
		client := proto.NewUsersServiceClient(con.Conn)
		res, err := client.Login(c, &proto.LoginRequest{Username: superSendInfo.Username, Password: superSendInfo.Password})
		if err != nil {
			ResponseFailed(c, nil, err.Error())
			return
		} else {
			if res.Code == 1 {
				dao.Update(c, "id=@id", map[string]interface{}{"token": res.LoginResponse.Token}, map[string]interface{}{"id": c.GetInt("super_send_id")})
				ResponseSuccess(c, res.LoginResponse.Token, res.Message)
			} else {
				ResponseFailed(c, nil, res.Message)
			}
			return
		}
	}
}

// @Summary 注册super_send
// @Router /super_send/registerSuperSend [POST]
func RegisterSuperSend(c *gin.Context) {
	dao := dbmodel.NewSuperSendConnInfoDao(c, db.Db)
	superSendInfo, err := dao.Get(c, "*", "id=@id", map[string]interface{}{"id": c.GetInt("super_send_id")})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
		if con == nil {
			ResponseFailed(c, nil, "连接失败")
			return
		}
		client := proto.NewUsersServiceClient(con.Conn)
		// 创建元数据
		md := metadata.Pairs(
			"token", c.GetString("token"),
		)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		ctx = metadata.NewOutgoingContext(ctx, md)
		res, err := client.Register(ctx, &proto.RegisterRequest{Username: superSendInfo.Username, Password: superSendInfo.Password})
		if err != nil {
			ResponseFailed(c, nil, err.Error())
			return
		} else {
			if res.Code == 1 {
				resLogin, _ := client.Login(c, &proto.LoginRequest{
					Username: superSendInfo.Username,
					Password: superSendInfo.Password,
				})
				dao.Update(c, "id=@id", map[string]interface{}{"token": resLogin.LoginResponse.Token}, map[string]interface{}{"id": c.GetInt("super_send_id")})
				ResponseSuccess(c, resLogin.LoginResponse.Token, res.Message)
			} else {
				ResponseFailed(c, nil, res.Message)
			}
			return
		}
	}
}
