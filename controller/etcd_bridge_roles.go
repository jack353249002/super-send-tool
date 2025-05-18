package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"super-send-tool/etcdbridgeproto"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

func AddEtcdBridgeRole(c *gin.Context) {
	var params request.AddRolesRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	actions := strings.Split(params.Action, ",")
	res, err := client.Add(ctx, &etcdbridgeproto.RolesRequest{Role: params.Role, Path: params.Path, Action: actions})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.GetCode() == 1 {
			ResponseSuccess(c, nil, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}

func GetEtcdBridgeRolesList(c *gin.Context) {
	var rolesRequest request.GetRolesRequest
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetList(ctx, &etcdbridgeproto.RolesLisRequest{SearchKey: rolesRequest.KeyWords, Page: int64(rolesRequest.PageNo), PageSize: int64(rolesRequest.PageSize)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		if res.GetCode() == 1 {
			countData := len(res.GetRolesList().GetData())
			totalPage := utils.CreatePage(countData, countData)
			ResponseTableSuccess(c, res.GetRolesList().GetData(), countData, 1, totalPage, countData, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func DelEtcdBridgeRole(c *gin.Context) {
	var params request.DelRoleRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	var rolesPermissions []*etcdbridgeproto.RolesPermissions
	for _, val := range params {
		row := &etcdbridgeproto.RolesPermissions{Role: val.Role, Path: val.Path, Action: val.Action}
		rolesPermissions = append(rolesPermissions, row)
	}
	res, err := client.DelRole(ctx, &etcdbridgeproto.RoleDelRequest{Roles: rolesPermissions})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func AddEtcdBridgeRoleForUser(c *gin.Context) {
	var params request.AddRoleForUserRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.AddRoleForUser(ctx, &etcdbridgeproto.AddRoleForUserRequest{Role: params.Role, Username: params.Username})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func DelEtcdBridgeRoleForUser(c *gin.Context) {
	var params request.DelRoleForUserRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.DelUserRole(ctx, &etcdbridgeproto.DelRoleForUserRequest{Role: params.Role, Username: params.Username})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func GetEtcdBridgeRolesPermissions(c *gin.Context) {
	var params request.GetRolesPermissionsRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	actions := strings.Split(params.Action, ",")
	res, err := client.GetRolesPermissions(ctx, &etcdbridgeproto.RolesRequest{Role: params.Role, Path: params.Path, Action: actions})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, res.GetRolesPermissionsList().GetData(), res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func DelEtcdBridgeRolesPermissions(c *gin.Context) {
	var params request.DelRolesPermissionsRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.EtcdBridgeErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	var rolesPermissions []*etcdbridgeproto.RolesPermissions
	for _, val := range params {
		row := &etcdbridgeproto.RolesPermissions{Role: val.Role, Path: val.Path, Action: val.Action}
		rolesPermissions = append(rolesPermissions, row)
	}
	res, err := client.DelRolesPermissions(ctx, &etcdbridgeproto.RolesPermissionsDelRequest{Roles: rolesPermissions})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				ResponseNotAuth(c, nil, statusErr.Message())
			// 其他状态码...
			default:
				ResponseFailed(c, nil, statusErr.Message())
			}
		}
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
