package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"io"
	"super-send-tool/grpccons"
	"super-send-tool/model/response"
	"super-send-tool/proto"
	"sync"
	"time"
)

var ReloadSignal = ReloadSignalModel{
	Token: "",
	Logs:  make([]*proto.SmtpResponse, 0),
}

// 重新加载服务信号
type ReloadSignalModel struct {
	Locker      sync.Mutex
	TokenLocker sync.Mutex
	Token       string //访问令牌
	Logs        []*proto.SmtpResponse
	IsEnd       bool
	Context     *gin.Context
	EndTime     int64 //当前令牌的结束时间
}

func (r *ReloadSignalModel) Start() {
	r.Locker.Lock()
}
func (r *ReloadSignalModel) End() {
	r.TokenLocker.Lock()
	if r.Token != "" {
		r.Token = ""
		r.Locker.Unlock()
		r.ClearLog()
	}
	r.TokenLocker.Unlock()
}
func (r *ReloadSignalModel) CreateToken(c *gin.Context) string {
	r.Start()
	r.Context = c
	newUUID := uuid.New()
	r.Token = newUUID.String()
	return r.Token
}
func (r *ReloadSignalModel) Reload() {
	r.IsEnd = false
	go func() {
		defer func() {
			oldToken := r.Token
			r.IsEnd = true
			r.EndTime = time.Now().UTC().Unix()
			time.Sleep(time.Second * 10)
			if r.Token == oldToken && r.Token != "" {
				r.End()
			}
		}()
		con := grpccons.SuperSendGroupAction.Get(r.Context.GetInt("super_send_id"))
		if con == nil {
			r.AddLog(&proto.SmtpResponse{Code: 0, Message: "连接失败"})
			return
		}
		client := proto.NewSmtpServiceClient(con.Conn)
		// 创建元数据
		md := metadata.Pairs(
			"token", r.Context.GetString("token"),
		)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		ctx = metadata.NewOutgoingContext(ctx, md)
		res, err := client.Reload(ctx, &empty.Empty{})
		if err != nil {
			r.AddLog(&proto.SmtpResponse{Code: 0, Message: err.Error()})
			return
		} else {
			for {
				resp, err := res.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					r.AddLog(&proto.SmtpResponse{Code: 0, Message: err.Error()})
					return
				}
				r.AddLog(resp)
				time.Sleep(5 * time.Second)
			}
		}
	}()
}
func (r *ReloadSignalModel) ClearLog() {
	r.Logs = make([]*proto.SmtpResponse, 0)
}
func (r *ReloadSignalModel) AddLog(log *proto.SmtpResponse) {
	r.Logs = append(r.Logs, log)
}

type TableResponse struct {
	Data       interface{} `json:"data""`
	PageSize   int         `json:"pageSize""`
	PageNo     int         `json:"pageNo"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
}

func ResponseTableSuccess(c *gin.Context, data interface{}, pageSize, pageNo, totalPage, totalCount int, msg string) {
	var tableResponse TableResponse
	tableResponse.PageSize = pageSize
	tableResponse.PageNo = pageNo
	tableResponse.TotalPage = totalPage
	tableResponse.TotalCount = totalCount
	tableResponse.Data = data
	var base response.Base
	base.Status = 200
	base.Result = &tableResponse
	base.Message = msg
	c.JSON(200, base)
}
func ResponseSuccess(c *gin.Context, data interface{}, msg string) {
	var base response.Base
	base.Status = 200
	base.Result = data
	base.Message = msg
	c.JSON(200, base)
}
func ResponseFailed(c *gin.Context, data interface{}, msg string) {
	var base response.Base
	base.Status = 500
	base.Result = data
	base.Message = msg
	c.JSON(200, base)
}
func ResponseNotAuth(c *gin.Context, data interface{}, msg string) {
	var base response.Base
	base.Status = 401
	base.Result = data
	base.Message = msg
	c.JSON(200, base)
}
