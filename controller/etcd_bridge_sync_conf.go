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

func AddEtcdBridgeSyncConf(c *gin.Context) {
	var params request.AddEtcdBridgeSyncConfRequest
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
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Add(ctx, &etcdbridgeproto.SaveSyncConfRequest{
		SourceServerId: params.SourceServerID,
		TargetServerId: params.TargetServerID,
		WatcherKey:     params.WatcherKey,
		SyncAllTime:    params.SyncAllTime,
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

func UpdateEtcdBridgeSyncConf(c *gin.Context) {
	var params request.UpdateEtcdBridgeSyncConfRequest
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
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Update(ctx, &etcdbridgeproto.SaveSyncConfRequest{
		Id:             int64(params.ID),
		SourceServerId: params.SourceServerID,
		TargetServerId: params.TargetServerID,
		WatcherKey:     params.WatcherKey,
		SyncAllTime:    params.SyncAllTime,
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

func DelEtcdBridgeSyncConf(c *gin.Context) {
	var params request.DelEtcdBridgeSyncConfRequest
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
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Del(ctx, &etcdbridgeproto.DelSyncConfRequest{
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
func GetEtcdBridgeSyncConf(c *gin.Context) {
	var params request.GetEtcdBridgeSyncConfRequest
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
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Get(ctx, &etcdbridgeproto.GetSyncConfRequest{
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
			ResponseSuccess(c, res.GetSyncConfInfo(), res.GetMsg())
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
func GetEtcdBridgeSyncConfList(c *gin.Context) {
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
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetList(ctx, &etcdbridgeproto.SyncConfListRequest{SearchKey: getRequest.KeyWords, Page: int64(getRequest.PageNo), PageSize: int64(getRequest.PageSize)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		syncConfList := res.GetSyncConfList()
		if res.Code == 1 {
			totalPage := utils.CreatePage(int(syncConfList.GetTotal()), getRequest.PageSize)
			ResponseTableSuccess(c, syncConfList.GetData(), int(getRequest.PageSize), 1, totalPage, int(syncConfList.GetTotal()), "success")
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
}

func EtcdBridgeSyncAction(c *gin.Context) {
	var params request.EtcdBridgeSyncConfActionRequest
	err := c.ShouldBind(&params)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	con := grpccons.EtcdBridgeGroupAction.Get(c.GetInt("etcd_bridge_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := etcdbridgeproto.NewSyncConfServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Action(ctx, &etcdbridgeproto.SyncConfActionRequest{Id: int64(params.ID), Status: int64(params.Status)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, "success")
		} else {
			ResponseFailed(c, nil, res.GetMsg())
		}
		return
	}
	return
}
