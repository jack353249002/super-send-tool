package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

func AddRole(c *gin.Context) {
	var params request.AddRolesRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	actions := strings.Split(params.Action, ",")
	res, err := client.Add(ctx, &proto.RolesRequest{Role: params.Role, Path: params.Path, Action: actions})
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
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

func GetRolesList(c *gin.Context) {
	var rolesRequest request.GetRolesRequest
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetList(ctx, &proto.RolesLisRequest{SearchKey: rolesRequest.KeyWords, Page: int64(rolesRequest.PageNo), PageSize: int64(rolesRequest.PageSize)})
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
			countData := len(res.Data)
			totalPage := utils.CreatePage(countData, countData)
			ResponseTableSuccess(c, res.GetData(), countData, 1, totalPage, countData, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func DelRole(c *gin.Context) {
	var params request.DelRoleRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	var rolesPermissions []*proto.RolesPermissions
	for _, val := range params {
		row := &proto.RolesPermissions{Role: val.Role, Path: val.Path, Action: val.Action}
		rolesPermissions = append(rolesPermissions, row)
	}
	res, err := client.DelRole(ctx, &proto.RoleDelRequest{Roles: rolesPermissions})
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
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func AddRoleForUser(c *gin.Context) {
	var params request.AddRoleForUserRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.AddRoleForUser(ctx, &proto.AddRoleForUserRequest{Role: params.Role, Username: params.Username})
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
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func DelRoleForUser(c *gin.Context) {
	var params request.DelRoleForUserRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.DelUserRole(ctx, &proto.DelRoleForUserRequest{Role: params.Role, Username: params.Username})
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
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func GetRolesPermissions(c *gin.Context) {
	var params request.GetRolesPermissionsRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	actions := strings.Split(params.Action, ",")
	res, err := client.GetRolesPermissions(ctx, &proto.RolesRequest{Role: params.Role, Path: params.Path, Action: actions})
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
			ResponseSuccess(c, res, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func DelRolesPermissions(c *gin.Context) {
	var params request.DelRolesPermissionsRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		// 自定义错误提示
		errors := validate.SuperSendErrorMessages(err)
		// 提取验证错误信息
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewRolesServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	var rolesPermissions []*proto.RolesPermissions
	for _, val := range params {
		row := &proto.RolesPermissions{Role: val.Role, Path: val.Path, Action: val.Action}
		rolesPermissions = append(rolesPermissions, row)
	}
	res, err := client.DelRolesPermissions(ctx, &proto.RolesPermissionsDelRequest{Roles: rolesPermissions})
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
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
