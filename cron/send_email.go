package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"super-send-tool/db"
	"super-send-tool/grpccons"
	"super-send-tool/model/dbmodel"
	"super-send-tool/proto"
	"time"
)

// 发送邮件定时任务
func SendEmailCron() {
	for {
		SendEmail()
		time.Sleep(60 * time.Second)
	}
}
func SendEmail() {
	nowTime := time.Now().UTC()
	nowTimeDayBegin := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.UTC)
	ctx := context.Background()
	dao := dbmodel.NewCheckUserAliveDao(ctx, db.Db)
	list, _ := dao.ListAll(ctx, "*", "@nowDayBegin - day_login_first_time >= send_email_action_timeout  AND last_send_email_time < @nowDayBegin2", map[string]interface{}{
		"nowDayBegin":  nowTimeDayBegin.Unix(),
		"nowDayBegin2": nowTimeDayBegin.Unix(),
	})
	for _, user := range list {
		var tempUser dbmodel.CheckUserAlive
		tempUser = user
		go func() {
			con := grpccons.SuperSendGroupAction.Get(user.SuperSendConnInfoId)
			if con == nil {
				return
			}
			ok := ResetSend(&tempUser)
			if ok {
				dao.Update(ctx, "id=@id", map[string]interface{}{"last_send_email_time": nowTime.Unix()}, map[string]interface{}{"id": tempUser.ID})
			}
			return
		}()
	}
}
func ResetSend(user *dbmodel.CheckUserAlive) bool {
	con := grpccons.SuperSendGroupAction.Get(user.SuperSendConnInfoId)
	if con == nil {
		return false
	}
	client := proto.NewSendServiceClient(con.Conn)
	// 创建元数据
	md := metadata.Pairs(
		"token", con.Token,
	)
	conCtx, cancel := context.WithTimeout(context.Background(), 115*time.Second)
	defer cancel()
	outCtx := metadata.NewOutgoingContext(conCtx, md)
	params := map[string]string{
		"position": user.Position,
	}
	paramsJson, _ := json.Marshal(params)
	title := fmt.Sprintf("[userID:%d,userName:%s]的紧急联系邮件发送", user.ID, user.Username)
	resp, err := client.SetSend(outCtx, &proto.AddSendRequest{Title: title, MessageId: int64(user.MessageID), Receive: user.SendEmailAccounts, Createtime: time.Now().UTC().Unix(), SendModel: 0,
		SendServerId: user.SmtpIds, DispatchModel: 0, Params: string(paramsJson)})
	var sendInfoResp *proto.SendInfoActionResponse
	if err == nil && resp.GetCode() == 1 {
		//resp, err = client.SetSendInfo(outCtx, &proto.EditSendRequest{Id: int64(user.SendID), Params: string(paramsJson)})
		sendInfoResp, err = client.SendInfoAction(outCtx, &proto.SendInfoActionRequest{SendId: resp.GetData().Id, Status: 1})
		if sendInfoResp.GetCode() == 1 {
			return true
		} else {
			return false
		}
	} else if err == nil && resp.GetCode() == 0 {
		return false
	} else {
		statusErr, ok := status.FromError(err)
		if ok {
			switch statusErr.Code() {
			case codes.Unauthenticated:
				superSendInfoDao := dbmodel.NewSuperSendConnInfoDao(conCtx, db.Db)
				userClient := proto.NewUsersServiceClient(con.Conn)
				loginRes, err := userClient.Login(conCtx, &proto.LoginRequest{Username: con.UserName, Password: con.Password})
				if err == nil && loginRes.Code == 1 {
					nowTime := time.Now().UTC().Unix()
					superSendInfoDao.Update(conCtx, "id=@id", map[string]interface{}{"token": loginRes.LoginResponse.Token, "conn_last_login_time": nowTime}, map[string]interface{}{"id": user.SuperSendConnInfoId})
					con.Token = loginRes.LoginResponse.Token
					md = metadata.Pairs(
						"token", con.Token,
					)
					conCtx, _ = context.WithTimeout(context.Background(), 115*time.Second)
					outCtx = metadata.NewOutgoingContext(conCtx, md)
					resp, err = client.SetSend(outCtx, &proto.AddSendRequest{Title: title, MessageId: int64(user.MessageID), Receive: user.SendEmailAccounts, Createtime: time.Now().UTC().Unix(), SendModel: 0,
						SendServerId: user.SmtpIds, DispatchModel: 0, Params: string(paramsJson)})
					if err == nil && resp.GetCode() == 1 {
						sendInfoResp, err = client.SendInfoAction(outCtx, &proto.SendInfoActionRequest{SendId: resp.GetData().Id, Status: 1})
						if sendInfoResp.GetCode() == 1 {
							return true
						} else {
							return false
						}
					} else {
						return false
					}
				}
			default:
				return false
			}
		}
		return false
	}
}
