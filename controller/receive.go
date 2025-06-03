package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	client2 "super-send-tool/client"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/model/response"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 添加接收规则
// @Router /super_send/addReceive [POST]
func AddReceive(c *gin.Context) {
	var receive request.ReceiveRequest
	err := c.ShouldBind(&receive)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.AddReceive(ctx, &proto.SetReceiveRequest{Title: receive.Title, ReceiveModel: int64(receive.ReceiveModel), ReceiveRule: receive.ReceiveRule, ReceiveServerId: receive.ReceiveServerID, EndType: int64(receive.EndType), EndRule: receive.EndRule, Status: int64(receive.Status)})
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

// @Summary 监听新邮件信息
// @Router /super_send/addNewReceiveEmailMessageListen [POST]
func AddNewReceiveEmailMessageListen(c *gin.Context) {
	var receiveEmailMessageListen request.ListenReceiveMessage
	err := c.ShouldBind(&receiveEmailMessageListen)
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
	receiveClient := client2.PublicClientPool.GetClient(c.GetInt("super_send_id"))
	if receiveClient != nil {
		receiveClient.ListenReceive(receiveEmailMessageListen.IDs, c.GetString("token"), con.UserName)
		ResponseSuccess(c, nil, "listen success")
		return
	} else {
		ResponseFailed(c, nil, "listen error")
		return
	}

}

// @Summary 设置接收规则
// @Router /super_send/setReceive [POST]
func SetReceive(c *gin.Context) {
	var receive request.ReceiveRequest
	err := c.ShouldBind(&receive)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SetReceive(ctx, &proto.SetReceiveRequest{Id: int64(receive.ID), ReceiveServerId: receive.ReceiveServerID})
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

// @Summary 获取接收列表
// @Router /super_send/getReceiveList [POST]
func GetReceiveList(c *gin.Context) {
	var receiveListRequest request.ReceiveListRequest
	err := c.ShouldBind(&receiveListRequest)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetReceiveList(ctx, &proto.ReceiveListRequest{Page: int64(receiveListRequest.PageNo), Limit: int64(receiveListRequest.PageSize), Status: int64(receiveListRequest.Status)})
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
			totalPage := utils.CreatePage(int(res.Count), receiveListRequest.PageSize)
			ResponseTableSuccess(c, res.Data, int(res.Count), 1, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 获取邮件信息列表
// @Router /super_send/getReceiveMessagesList [POST]
func GetReceiveMessagesList(c *gin.Context) {
	var receiveListRequest request.ReceiveMessagesListRequest
	err := c.ShouldBind(&receiveListRequest)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetReceiveMessages(ctx, &proto.GetReceiveMessageListRequest{ReceiveId: int64(receiveListRequest.ReceiveID)})
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
			ResponseTableSuccess(c, res.Data, int(res.Count), 1, 1, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 获取邮件信息
// @Router /super_send/getReceiveMessageInfo [POST]
func GetReceiveMessagesInfo(c *gin.Context) {
	var receiveInfoRequest request.ReceiveMessagesInfoRequest
	err := c.ShouldBind(&receiveInfoRequest)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var emailInfo response.EmailInfo
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetReceiveMessageInfo(ctx, &proto.GetReceiveMessageInfoRequest{Id: int64(receiveInfoRequest.ID), ImapServerId: int64(receiveInfoRequest.ImapServerId), Uid: receiveInfoRequest.UID})
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
			emailInfo.MessageId = res.Data.MessageId
			emailInfo.SenderAddress = res.Data.SenderAddress
			emailInfo.Date = res.Data.Date
			emailInfo.Subject = res.Data.Subject
			emailInfo.ImapServerId = res.Data.ImapServerId
			emailInfo.Body = res.Data.Body
			emailInfo.ImapServerText = res.Data.ImapServerText
			if res.Data.Files != nil {
				dir := fmt.Sprintf("%d-%s-%s-%s", res.Data.ImapServerId, res.Data.Subject, res.Data.Date, res.Data.MessageId)
				_, hashed, err := utils.CreateFolderByMD5("run_root/filecache/", dir)
				if err == nil {
					for _, v := range res.Data.Files {
						if v.FileBytes != nil {
							fileName := fmt.Sprintf("run_root/filecache/%s/%s", hashed, v.FileName)
							utils.WriteBytesToFile(fileName, v.FileBytes)
							emailInfo.FilePaths = append(emailInfo.FilePaths, "/"+"filecache/"+hashed+"/"+v.FileName)
						}
					}
				}
			}
			ResponseSuccess(c, emailInfo, "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 消息接收指令行为
// @Router /super_send/receiveAction [POST]
func ReceiveAction(c *gin.Context) {
	var receiveActionRequest request.ReceiveMessageActionRequest
	err := c.ShouldBind(&receiveActionRequest)
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
	client := proto.NewReceiveServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.ReceiveAction(ctx, &proto.ReceiveActionRequest{Status: int64(receiveActionRequest.Status), ReceiveId: int64(receiveActionRequest.ReceiveID)})
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
