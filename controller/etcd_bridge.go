package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"super-send-tool/db"
	"super-send-tool/etcdbridgeproto"
	"super-send-tool/grpccons"
	"super-send-tool/model/dbmodel"
	"super-send-tool/model/request"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 获取列表
// @Router /etcd_bridge/getEtcdBridgeList [POST]
func GetEtcdBridgeList(c *gin.Context) {
	dao := dbmodel.NewEtcdBridgeConnInfoDao(c, db.Db)
	var pageListRequest request.EtcdBridgeListRequest
	err := c.ShouldBind(&pageListRequest)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
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

// @Summary 添加etcd_bridge
// @Router /etcd_bridge/addEtcdBridge [POST]
func AddEtcdBridge(c *gin.Context) {
	var addRequest request.AddEtcdBridgeRequest
	err := c.ShouldBind(&addRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewEtcdBridgeConnInfoDao(c, db.Db)
	err = dao.Create(c, &dbmodel.EtcdBridgeConnInfo{Address: addRequest.Address, IsSsl: int(addRequest.IsSSL), Password: addRequest.Password, Username: addRequest.UserName})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 编辑etcd_bridge
// @Router /etcd_bridge/updateEtcdBridge [POST]
func UpdateEtcdBridge(c *gin.Context) {
	var updateRequest request.UpdateEtcdBridgeRequest
	err := c.ShouldBind(&updateRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewEtcdBridgeConnInfoDao(c, db.Db)
	err = dao.Update(c, "id=@id", map[string]interface{}{"address": updateRequest.Address, "is_ssl": int(updateRequest.IsSSL), "password": updateRequest.Password, "username": updateRequest.UserName}, map[string]interface{}{"id": updateRequest.ID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 删除etcd_bridge
// @Router /etcd_bridge/deleteEtcdBridge [POST]
func DeleteEtcdBridge(c *gin.Context) {
	var deleteRequest request.DelEtcdBridgeRequest
	err := c.ShouldBind(&deleteRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	dao := dbmodel.NewEtcdBridgeConnInfoDao(c, db.Db)
	err = dao.Delete(c, "id=@id", map[string]interface{}{"id": deleteRequest.ID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	}
	ResponseSuccess(c, nil, "success")
	return
}

// @Summary 登录etcd_bridge
// @Router /etcd_bridge/loginEtcdBridge [POST]
func LoginEtcdBridge(c *gin.Context) {
	dao := dbmodel.NewEtcdBridgeConnInfoDao(c, db.Db)
	etcdBridgeInfo, err := dao.Get(c, "*", "id=@id", map[string]interface{}{"id": c.GetInt("etcd_bridge_id")})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
		if con == nil {
			ResponseFailed(c, nil, "连接失败")
			return
		}
		client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
		res, err := client.Login(c, &etcdbridgeproto.LoginRequest{Username: etcdBridgeInfo.Username, Password: etcdBridgeInfo.Password})
		if err != nil {
			ResponseFailed(c, nil, err.Error())
			return
		} else {
			if res.Code == 1 {
				loginResponse := res.Payload.(*etcdbridgeproto.BaseResponse_LoginInfo)
				dao.Update(c, "id=@id", map[string]interface{}{"token": loginResponse.LoginInfo.GetToken()}, map[string]interface{}{"id": c.GetInt("super_send_id")})
				ResponseSuccess(c, loginResponse.LoginInfo.GetToken(), res.GetMsg())
			} else {
				ResponseFailed(c, nil, res.GetMsg())
			}
			return
		}
	}
}

// @Summary 添加etcd_bridge用户
// @Router /etcd_bridge/addUser [POST]
func AddEtcdBridgeUser(c *gin.Context) {
	var addRequest request.AddEtcdBridgeUserRequest
	err := c.ShouldBind(&addRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Add(ctx, &etcdbridgeproto.UsersRequest{Username: addRequest.UserName, Password: addRequest.Password, Roles: addRequest.Roles, Status: addRequest.Status})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseSuccess(c, nil, res.GetMsg())
		return
	}
}

// @Summary 编辑etcd_bridge用户
// @Router /etcd_bridge/updateUser [POST]
func UpdateEtcdBridgeUser(c *gin.Context) {
	var updateRequest request.UpdateEtcdBridgeUserRequest
	err := c.ShouldBind(&updateRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Update(ctx, &etcdbridgeproto.UsersRequest{Id: int64(updateRequest.ID), Password: updateRequest.Password, Roles: updateRequest.Roles, Status: updateRequest.Status})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseSuccess(c, nil, res.GetMsg())
		return
	}
}

// @Summary 删除etcd_bridge用户
// @Router /etcd_bridge/delUser [POST]
func DelEtcdBridgeUser(c *gin.Context) {
	var delRequest request.DelEtcdBridgeRequest
	err := c.ShouldBind(&delRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Del(ctx, &etcdbridgeproto.DelUsersRequest{Id: int64(delRequest.ID)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseSuccess(c, nil, res.GetMsg())
		return
	}
}

// @Summary 获取etcd_bridge用户详情
// @Router /etcd_bridge/getUser [POST]
func GetEtcdBridgeUser(c *gin.Context) {
	var getRequest request.GetEtcdBridgeUserRequest
	err := c.ShouldBind(&getRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Get(ctx, &etcdbridgeproto.UsersRequest{Id: int64(getRequest.ID)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		ResponseSuccess(c, res.GetUsersInfo(), res.GetMsg())
		return
	}
}

// @Summary 获取etcd_bridge用户列表
// @Router /etcd_bridge/getUserList [POST]
func GetEtcdBridgeUserList(c *gin.Context) {
	var getRequest request.EtcdBridgeUsersRequest
	err := c.ShouldBind(&getRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.EtcdBridgeErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewUsersServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetList(ctx, &etcdbridgeproto.UsersLisRequest{SearchKey: getRequest.KeyWords, Page: int64(getRequest.PageNo), PageSize: int64(getRequest.PageSize)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		userList := res.GetUsersList()
		if res.Code == 1 {
			totalPage := utils.CreatePage(int(userList.GetTotal()), getRequest.PageSize)
			ResponseTableSuccess(c, userList.GetData(), int(getRequest.PageSize), 1, totalPage, int(userList.GetTotal()), "success")
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
}
