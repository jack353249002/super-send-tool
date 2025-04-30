package controller

import (
	"github.com/gin-gonic/gin"
	"super-send-tool/db"
	"super-send-tool/etcdbridgeproto"
	"super-send-tool/grpccons"
	"super-send-tool/model/dbmodel"
	"super-send-tool/model/request"
	"super-send-tool/validate"
)

// @Summary 添加etcd_bridge
// @Router /etcd_bridge/addEtcdBridge [POST]
func AddEtcdBridge(c *gin.Context) {
	var addRequest request.AddEtcdBridgeRequest
	err := c.ShouldBind(&addRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
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
		errors := validate.SuperSendErrorMessages(err)
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
		errors := validate.SuperSendErrorMessages(err)
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
