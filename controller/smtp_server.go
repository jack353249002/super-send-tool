package controller

import (
	"context"
	"github.com/gin-gonic/gin"
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

func GetSmtpServer(c *gin.Context) {
	var smtpServer request.SmtpListRequest
	err := c.ShouldBind(&smtpServer)
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
	client := proto.NewSmtpServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.GetSmtpList(ctx, &proto.GetSmtpListRequest{Keywords: smtpServer.KeyWords, Page: int64(smtpServer.PageNo), Limit: int64(smtpServer.PageSize)})
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
			totalPage := utils.CreatePage(int(res.Count), smtpServer.PageSize)
			ResponseTableSuccess(c, res.Data, smtpServer.PageSize, smtpServer.PageNo, totalPage, int(res.Count), "success")
		} else {
			ResponseFailed(c, nil, res.Message)
		}
		return
	}
	return
}
func SetSmtpServer(c *gin.Context) {
	var setSmtpRequest request.SetSmtpRequest
	err := c.ShouldBind(&setSmtpRequest)
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
	client := proto.NewSmtpServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.SetSmtp(ctx, &proto.SetSmtpRequest{Id: setSmtpRequest.ID, Title: setSmtpRequest.Title, SmtpServer: setSmtpRequest.SmtpServer, MaxConcurrency: setSmtpRequest.MaxConcurrency, Port: setSmtpRequest.Port, SmtpSendEmail: setSmtpRequest.SmtpSendEmail, SmtpPassword: setSmtpRequest.SmtpPassword})
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
func DelSmtpServer(c *gin.Context) {
	var delRequest request.DelSmtpRequest
	err := c.ShouldBind(&delRequest)
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
	client := proto.NewSmtpServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", c.GetString("token"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.DelSmtp(ctx, &proto.DelSmtpRequest{Id: delRequest.ID})
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

func ReloadServer(c *gin.Context) {
	var reloadRequest request.ReloadRequest
	err := c.ShouldBind(&reloadRequest)
	if err != nil {
		// 提取验证错误信息
		errors := validate.SuperSendErrorMessages(err)
		ResponseFailed(c, nil, validate.GetAllErrorStr(errors))
		return
	}
	if reloadRequest.Token == "" {
		token := ReloadSignal.CreateToken(c)
		ReloadSignal.Reload()
		ResponseSuccess(c, token, "success")
	} else if reloadRequest.Token == ReloadSignal.Token {
		if ReloadSignal.IsEnd {
			ResponseSuccess(c, ReloadSignal.Logs, "success")
			ReloadSignal.End()
			return
		} else {
			ResponseSuccess(c, ReloadSignal.Logs, "success")
			return
		}
	} else {
		ResponseFailed(c, nil, "token错误")
	}
	return
}
