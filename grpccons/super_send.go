package grpccons

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"super-send-tool/config/baseconfig"
	"super-send-tool/db"
	"super-send-tool/model/dbmodel"
	"sync"
	"time"
)

var SuperSendGroupAction = new(SuperSendGroup)

type SuperSendGroup struct {
	SuperSendMap sync.Map
}

// 获取发送器
func (s *SuperSendGroup) Get(id int) (superSend *SuperSend) {
	ctx := context.Background()
	superSend = s.GetSuperSend(id)
	if superSend == nil {
		dao := dbmodel.NewSuperSendConnInfoDao(ctx, db.Db)
		superSendInfo, err := dao.Get(ctx, "*", "id=@id", map[string]interface{}{"id": id})
		if err == nil {
			superSend = new(SuperSend)
			isSSL := false
			if superSendInfo.IsSsl == 1 {
				isSSL = true
			}
			err = superSend.Init(superSendInfo.ID, superSendInfo.Address, superSendInfo.Username, superSendInfo.Password, isSSL)
			if err == nil {
				s.SetSuperSend(superSend)
			} else {
				return nil
			}
		}
	}
	return
}

// 关闭发送器
func (s *SuperSendGroup) Close(id int) bool {
	superSend := s.GetSuperSend(id)
	if superSend != nil {
		s.SuperSendMap.Delete(id)
		superSend.Close()
		return true
	} else {
		return false
	}
}

// 根据id设置发送器
func (s *SuperSendGroup) SetSuperSend(superSend *SuperSend) {
	s.SuperSendMap.Store(superSend.ID, superSend)
}

// 根据地址获取发送器
func (s *SuperSendGroup) GetSuperSend(id int) (superSend *SuperSend) {
	val, have := s.SuperSendMap.Load(id)
	if have {
		superSend = val.(*SuperSend)
	} else {
		superSend = nil
	}
	return
}

// 用于控制每个发送器的 grpc 连接
type SuperSend struct {
	ID       int              //id
	Address  string           //连接地址
	Conn     *grpc.ClientConn //连接实体
	UserName string
	Password string
	IsSSL    bool
	Token    string
}

func (s *SuperSend) Init(id int, address string, userName, password string, isSSL bool) (err error) {
	s.ID = id
	s.Address = address
	s.UserName = userName
	s.Password = password
	s.IsSSL = isSSL
	var opts []grpc.DialOption
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	if s.IsSSL {
		caCertPool := x509.NewCertPool()
		for _, val := range baseconfig.CONFIG.CAPATHS {
			// 加载CA证书
			caCert, err := ioutil.ReadFile(val)
			if err != nil {
				continue
			}
			caCertPool.AppendCertsFromPEM(caCert)
		}
		// 创建TLS配置
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}
		// 创建gRPC凭证
		creds := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())
	kp := keepalive.ClientParameters{
		Time:                10 * time.Second, // 每10秒ping一次
		Timeout:             20 * time.Second, // 等待ack最多5秒
		PermitWithoutStream: true,             // 即使没有活跃流也允许ping
	}
	opts = append(opts, grpc.WithKeepaliveParams(kp))
	s.Conn, err = grpc.DialContext(ctx, address, opts...)
	fmt.Println(err)
	return
}
func (s *SuperSend) ReConnect() (err error) {
	if s.Conn != nil {
		s.Close()
	}
	s.Conn, err = grpc.Dial(s.Address, grpc.WithInsecure(), grpc.WithBlock())
	return
}
func (s *SuperSend) Close() (err error) {
	s.Conn.Close()
	s.Conn = nil
	return
}
