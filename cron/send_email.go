package cron

import (
	"context"
	"encoding/json"
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
	nowTime := time.Now()
	nowTimeDayBegin := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
	ctx := context.Background()
	dao := dbmodel.NewCheckUserAliveDao(ctx, db.Db)
	list, _ := dao.ListAll(ctx, "*", "@nowDayBegin - day_login_first_time >= send_email_action_timeout  AND last_send_email_time < @nowDayBegin2", map[string]interface{}{
		"nowDayBegin":  nowTimeDayBegin.Unix(),
		"nowDayBegin2": nowTimeDayBegin.Unix(),
	})
	for _, user := range list {
		con := grpccons.SuperSendGroupAction.Get(user.SuperSendConnInfoId)
		if con == nil {
			continue
		}
		client := proto.NewSendServiceClient(con.Conn)
		ok := ResetSend(&user)
		if ok {
			// 创建元数据
			md := metadata.Pairs(
				"token", con.Token,
			)
			conCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			outCtx := metadata.NewOutgoingContext(conCtx, md)
			resp, err := client.SendInfoAction(outCtx, &proto.SendInfoActionRequest{SendId: int64(user.SendID), Status: 1})
			if err != nil || resp.GetCode() != 1 {
				continue
			} else {
				dao.Update(ctx, "id=@id", map[string]interface{}{"last_send_email_time": nowTimeDayBegin.Unix()}, map[string]interface{}{"id": user.ID})
			}
		}
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
	resp, err := client.ResetSend(outCtx, &proto.EditSendRequest{Id: int64(user.SendID)})
	if err == nil && resp.GetCode() == 1 {
		resp, err = client.SetSendInfo(outCtx, &proto.EditSendRequest{Id: int64(user.SendID), Params: string(paramsJson)})
		return true
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
					nowTime := time.Now().Unix()
					superSendInfoDao.Update(conCtx, "id=@id", map[string]interface{}{"token": loginRes.LoginResponse.Token, "conn_last_login_time": nowTime}, map[string]interface{}{"id": user.SuperSendConnInfoId})
					con.Token = loginRes.LoginResponse.Token
					md = metadata.Pairs(
						"token", con.Token,
					)
					conCtx, _ = context.WithTimeout(context.Background(), 115*time.Second)
					outCtx = metadata.NewOutgoingContext(conCtx, md)
					resp, err = client.ResetSend(outCtx, &proto.EditSendRequest{Id: int64(user.SendID)})
					if err == nil && resp.GetCode() == 1 {
						resp, err = client.SetSendInfo(outCtx, &proto.EditSendRequest{Id: int64(user.SendID), Params: string(paramsJson)})
						return true
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
