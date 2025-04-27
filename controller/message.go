package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"super-send-tool/grpccons"
	"super-send-tool/model/request"
	"super-send-tool/proto"
	"super-send-tool/utils"
	"super-send-tool/validate"
	"time"
)

// @Summary 文件上传
// @Router /super_send/upload [POST]
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ResponseFailed(c, nil, "文件上传失败: "+err.Error())
		return
	}
	// 打开文件
	src, err := file.Open()
	if err != nil {
		ResponseFailed(c, nil, "文件打开失败: "+err.Error())
		return
	}
	defer src.Close()

	// 读取文件内容到字节切片
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		ResponseFailed(c, nil, "读取文件失败: "+err.Error())
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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.UploadFile(ctx, &proto.FileRequest{FileName: file.Filename, Data: fileBytes})
	if err != nil {
		ResponseFailed(c, nil, err.Error())
		return
	} else {
		if res.Code == 1 {
			ResponseSuccess(c, map[string]interface{}{"url": ".exe", "path": res.GetData()}, res.Message)
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}

// @Summary 设置消息
// @Router /super_send/setMessage [POST]
func SetMessage(c *gin.Context) {
	var setMessageRequest request.SetMessageRequest
	err := c.ShouldBind(&setMessageRequest)
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
	res, err := client.SetMessage(ctx, &proto.MessageRequest{Id: int64(setMessageRequest.ID), Title: setMessageRequest.Title, Body: setMessageRequest.Body, Attach: setMessageRequest.Attach, ContentType: setMessageRequest.ContentType})
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

// @Summary 删除消息
// @Router /super_send/delMessage [POST]
func DelMessage(c *gin.Context) {
	var delMessageRequest request.DelMessageRequest
	err := c.ShouldBind(&delMessageRequest)
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
	res, err := client.DelMessage(ctx, &proto.MessageRequest{Id: int64(delMessageRequest.ID)})
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
			totalPage := utils.CreatePage(int(res.Count), pageListRequest.PageSize)
			ResponseTableSuccess(c, res.Data, pageListRequest.PageSize, pageListRequest.PageNo, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
