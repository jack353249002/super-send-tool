package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 创建消息
// @Router /super_send/setMessage [POST]
func SetMessage(c *gin.Context) {
	var addMessageRequest request.AddMessageRequest
	err := c.ShouldBind(&addMessageRequest)
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
	client := proto.NewMessageServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SetMessage(ctx, &proto.MessageRequest{Title: addMessageRequest.Title, Body: addMessageRequest.Body, Attach: addMessageRequest.Attach})
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

// @Summary 获取列表
// @Router /super_send/getMessageList [POST]
func GetMessageList(c *gin.Context) {
	var pageListRequest request.MessageListRequest
	err := c.ShouldBind(&pageListRequest)
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
	client := proto.NewMessageServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetMessageList(ctx, &proto.MessageListRequest{Title: pageListRequest.KeyWords, Page: int64(pageListRequest.PageNo), Limit: int64(pageListRequest.PageSize), StartTime: pageListRequest.StartTime, EndTime: pageListRequest.EndTime})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		if res.Code == 1 {
			totalPage := utils.CreatePage(int(res.Count), pageListRequest.PageSize)
			ResponseTableSuccess(c, res.Data, pageListRequest.PageSize, pageListRequest.PageNo, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
