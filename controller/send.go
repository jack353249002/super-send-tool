package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 设置发送配置
// @Router /super_send/setSend [POST]
func SetSend(c *gin.Context) {
	var setSendRequest request.SetSendRequest
	err := c.ShouldBind(&setSendRequest)
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
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SetSend(ctx, &proto.AddSendRequest{Title: setSendRequest.Title, Receive: setSendRequest.Receive, MessageId: setSendRequest.MessageID,
		SendModel: setSendRequest.SendModel, DispatchModel: setSendRequest.DispatchModel, SendServerId: setSendRequest.SendServerID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, nil, res.Message)
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 获取发送配置列表
// @Router /super_send/getSendInfoList [POST]
func GetSendInfoList(c *gin.Context) {
	var getSendInfoListRequest request.GetSendInfoListRequest
	err := c.ShouldBind(&getSendInfoListRequest)
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
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetSendInfoList(ctx, &proto.GetSendInfoListRequest{Page: int64(getSendInfoListRequest.PageNo), Limit: int64(getSendInfoListRequest.PageSize), Status: int64(getSendInfoListRequest.Status)})
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
			totalPage := utils.CreatePage(int(res.Count), getSendInfoListRequest.PageSize)
			ResponseTableSuccess(c, res.GetData(), getSendInfoListRequest.PageSize, getSendInfoListRequest.PageNo, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 获取发送接收者配置列表
// @Router /super_send/getSendList [POST]
func GetSendList(c *gin.Context) {
	var getSendListRequest request.GetSendListRequest
	err := c.ShouldBind(&getSendListRequest)
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
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetSendList(ctx, &proto.GetSendListRequest{Page: int64(getSendListRequest.PageNo), Limit: int64(getSendListRequest.PageSize), Status: int64(getSendListRequest.Status), SendId: int64(getSendListRequest.SendID)})
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
			totalPage := utils.CreatePage(int(res.Count), getSendListRequest.PageSize)
			ResponseTableSuccess(c, res.GetData(), getSendListRequest.PageSize, getSendListRequest.PageNo, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 消息发送指令行为
// @Router /super_send/sendInfoAction [POST]
func SendInfoAction(c *gin.Context) {
	var sendActiontRequest request.SendInfoActionRequest
	err := c.ShouldBind(&sendActiontRequest)
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
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SendInfoAction(ctx, &proto.SendInfoActionRequest{SendId: int64(sendActiontRequest.SendID), Status: int64(sendActiontRequest.Status)})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
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

// @Summary 编辑已经存在的发送配置信息
// @Router /super_send/setSendInfo [POST]
func SetSendInfo(c *gin.Context) {
	var setSendInfoRequest request.SetSendInfoRequest
	err := c.ShouldBind(&setSendInfoRequest)
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
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SetSendInfo(ctx, &proto.EditSendRequest{Id: setSendInfoRequest.ID, SendServerId: setSendInfoRequest.SendServerID})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
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

// @Summary 获取发送接收者配置列表
// @Router /super_send/getSmtpServerList [POST]
func GetSmtpServerList(c *gin.Context) {
	con := grpccons.SuperSendGroupAction.Get(c.GetInt("super_send_id"))
	if con == nil {
		ResponseFailed(c, nil, "连接失败")
		return
	}
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetSmtpServerList(ctx, &empty.Empty{})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
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
