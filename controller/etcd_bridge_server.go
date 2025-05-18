package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"super-send-tool/etcdbridgeproto"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

func AddEtcdBridgeServer(c *gin.Context) {
	var params request.AddEtcdBridgeServerRequest
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
	client := etcdbridgeproto.NewServerInfoServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Add(ctx, &etcdbridgeproto.SaveServerInfoRequest{
		IsSsl: int32(params.IsSSL),
		Host:  params.Host, CaCerPath: params.CaCerPath,
		CertPath: params.CertPath,
		KeyPath:  params.KeyPath,
		Password: params.Password,
		Type:     int64(params.Type),
		Username: params.UserName,
	})
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

func UpdateEtcdBridgeServer(c *gin.Context) {
	var params request.UpdateEtcdBridgeServerRequest
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
	client := etcdbridgeproto.NewServerInfoServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Update(ctx, &etcdbridgeproto.SaveServerInfoRequest{
		Id:        int64(params.ID),
		IsSsl:     int32(params.IsSSL),
		Host:      params.Host,
		CaCerPath: params.CaCerPath,
		CertPath:  params.CertPath,
		KeyPath:   params.KeyPath,
		Password:  params.Password,
		Type:      int64(params.Type),
		Username:  params.UserName,
	})
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

func DelEtcdBridgeServer(c *gin.Context) {
	var params request.DelEtcdBridgeServerRequest
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
	client := etcdbridgeproto.NewServerInfoServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Del(ctx, &etcdbridgeproto.DelSaveServerInfoRequest{
		Id: int64(params.ID),
	})
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
func GetEtcdBridgeServer(c *gin.Context) {
	var params request.GetEtcdBridgeServerRequest
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
	client := etcdbridgeproto.NewServerInfoServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Get(ctx, &etcdbridgeproto.GetServerInfoRequest{
		Id: int64(params.ID),
	})
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
			ResponseSuccess(c, res.GetServerInfo(), res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func GetEtcdBridgeServerList(c *gin.Context) {
	var getRequest request.PageListRequest
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
	client := etcdbridgeproto.NewServerInfoServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetList(ctx, &etcdbridgeproto.ServerInfoLisRequest{SearchKey: getRequest.KeyWords, Page: int64(getRequest.PageNo), PageSize: int64(getRequest.PageSize)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		serverInfoList := res.GetServerInfoList()
		if res.Code == 1 {
			totalPage := utils.CreatePage(int(serverInfoList.GetTotal()), getRequest.PageSize)
			ResponseTableSuccess(c, serverInfoList.GetData(), int(getRequest.PageSize), 1, totalPage, int(serverInfoList.GetTotal()), "success")
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
}
