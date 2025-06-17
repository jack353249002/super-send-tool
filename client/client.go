package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"super-send-tool/grpccons"
	"super-send-tool/proto"
	"sync"
	"time"
)

var PublicClientPool = new(ClientPool)

type ClientPool struct {
	Pool sync.Map
}

func (c *ClientPool) AddClient(id int, client *Client) {
	c.Pool.Store(id, client)
}
func (c *ClientPool) GetClient(id int) *Client {
	client, have := c.Pool.Load(id)
	if have {
		return client.(*Client)
	} else {
		return nil
	}
}

type UserCacheInfo struct {
	Username         string
	Password         string
	ReceiveIDsLocker sync.RWMutex
	ReceiveIDs       []int64
}
type ClientSignalInfo struct {
	Code int
	Data interface{}
}
type ClientSignal chan ClientSignalInfo
type Client struct {
	SuperSend         *grpccons.SuperSend
	UsersServerClient proto.UsersServiceClient
	Token             string
	Stream            proto.UsersService_MessageSendClient
	UserInfoCache     sync.Map
}

func (c *Client) SetUserInfoCache(username, password string, receiveIDs []int64) {
	var user UserCacheInfo
	user.Username = username
	user.Password = password
	user.ReceiveIDs = receiveIDs
	c.UserInfoCache.Store(username, &user)
}
func (c *Client) GetUserInfoCache(username string) *UserCacheInfo {
	res, ok := c.UserInfoCache.Load(username)
	if ok {
		return res.(*UserCacheInfo)
	} else {
		return nil
	}
}
func (c *Client) GetUserInfoCacheAll() (res []*UserCacheInfo) {
	c.UserInfoCache.Range(func(key, value any) bool {
		res = append(res, value.(*UserCacheInfo))
		return true
	})
	return
}
func (c *Client) Init(con *grpccons.SuperSend) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()
	c.SuperSend = con
	c.UsersServerClient = proto.NewUsersServiceClient(c.SuperSend.Conn)
	ctx := context.Background()
	c.Stream, err = c.UsersServerClient.MessageSend(ctx)
	return err
}
func (c *Client) Login(username, password string) error {
	messageSend := &proto.MessageSendRequest{Route: "login"}
	messageSend.Payload = &proto.MessageSendRequest_LoginRequest{
		LoginRequest: &proto.LoginRequest{
			Username: username,
			Password: password,
		},
	}
	err := c.Stream.Send(messageSend)
	user := c.GetUserInfoCache(username)
	if user == nil {
		c.SetUserInfoCache(username, password, []int64{})
	}
	fmt.Println("Login", err)
	return err
}
func (c *Client) ListenReceive(ids []int64, token string, username string) error {
	messageSend := &proto.MessageSendRequest{Route: "listenReceive", Token: token}
	messageSend.Payload = &proto.MessageSendRequest_ListenReceiveRequest{
		ListenReceiveRequest: &proto.ListenReceiveRequest{
			Ids: ids,
		},
	}
	err := c.Stream.Send(messageSend)
	user := c.GetUserInfoCache(username)
	fmt.Println("listenReceiveInfo", ids, user)
	if user != nil {
		user.ReceiveIDsLocker.Lock()
		user.ReceiveIDs = ids
		user.ReceiveIDsLocker.Unlock()
	}
	fmt.Println("listenReceiveErr", err)
	return err
}
func (c *Client) listenMessageAction(signal ClientSignal) {
	for {
		resp, err := c.Stream.Recv()
		fmt.Println("show", resp, err)
		if err == io.EOF {
			log.Println("Stream closed by server")
			var signInfo ClientSignalInfo
			signInfo.Code = -1
			signInfo.Data = nil
			signal <- signInfo
			break
		}
		if err != nil {
			log.Printf("Receive error: %v", err)
			var signInfo ClientSignalInfo
			signInfo.Code = -2
			signInfo.Data = nil
			signal <- signInfo
			break
		}
		switch resp.GetRoute() {
		case "login":
			log.Println("Login response:", resp.GetRoute())
			if resp.GetCode() == 1 {
				var signInfo ClientSignalInfo
				signInfo.Code = 1
				signInfo.Data = resp.GetLoginResponse().GetToken()
				signal <- signInfo
			}
		case "listenReceive":
			log.Println("listenReceive", resp.GetMessage())
			var signInfo ClientSignalInfo
			signInfo.Code = 2
			signInfo.Data = nil
			signal <- signInfo
		case "messageSend":
			log.Println("messageSend", resp.GetMessageInfo().GetSubject())
			var signInfo ClientSignalInfo
			signInfo.Code = 3
			signInfo.Data = resp.GetMessageInfo().GetSubject()
			signal <- signInfo
		}
		log.Printf("Received: %s", resp.Message)
	}
}
func (c *Client) ListenMessage() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	var signal ClientSignal
	signal = make(ClientSignal)
	autoLogin := false
	var autoLoginUserInfo *UserCacheInfo
	go c.listenMessageAction(signal)
	for waitLogin := range signal {
		if waitLogin.Code == 1 {
			if autoLogin {
				autoLogin = false
				c.ListenReceive(autoLoginUserInfo.ReceiveIDs, waitLogin.Data.(string), autoLoginUserInfo.Username)
			}
		}
		if waitLogin.Code == -2 || waitLogin.Code == -1 {
			grpccons.SuperSendGroupAction.Close(c.SuperSend.ID)
			con := grpccons.SuperSendGroupAction.Get(c.SuperSend.ID)
			err := c.Init(con)
			if err != nil {
				for {
					err = c.Init(con)
					if err == nil {
						break
					}
					fmt.Println("initerror", err)
					time.Sleep(5 * time.Second)
				}
			}
			go c.listenMessageAction(signal)
			users := c.GetUserInfoCacheAll()
			for _, user := range users {
				fmt.Println("logins", user)
				autoLoginUserInfo = c.GetUserInfoCache(user.Username)
				if autoLoginUserInfo == nil {
					autoLoginUserInfo.Username = user.Username
					autoLoginUserInfo.Password = user.Password
					autoLoginUserInfo.ReceiveIDs = user.ReceiveIDs
				}
				c.Login(user.Username, user.Password)
				autoLogin = true
			}
		}
	}
}
