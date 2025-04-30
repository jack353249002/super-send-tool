package grpccons

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"super-send-tool/config/baseconfig"
	"super-send-tool/db"
	"super-send-tool/model/dbmodel"
	"sync"
	"time"
)

var EtcdBridgeGroupAction = new(EtcdBridgeGroup)

type EtcdBridgeGroup struct {
	EtcdBridgeMap sync.Map
}

// 获取发送器
func (s *EtcdBridgeGroup) Get(id int) (etcdBridge *EtcdBridge) {
	ctx := context.Background()
	etcdBridge = s.GetEtcdBridge(id)
	if etcdBridge == nil {
		dao := dbmodel.NewEtcdBridgeConnInfoDao(ctx, db.Db)
		etcdBridgeInfo, err := dao.Get(ctx, "*", "id=@id", map[string]interface{}{"id": id})
		if err == nil {
			etcdBridge = new(EtcdBridge)
			isSSL := false
			if etcdBridgeInfo.IsSsl == 1 {
				isSSL = true
			}
			err = etcdBridge.Init(etcdBridgeInfo.ID, etcdBridgeInfo.Address, etcdBridgeInfo.Username, etcdBridgeInfo.Password, isSSL)
			if err == nil {
				s.SetEtcdBridge(etcdBridge)
			} else {
				return nil
			}
		}
	}
	return
}

// 关闭发送器
func (s *EtcdBridgeGroup) Close(id int) bool {
	etcdBridge := s.GetEtcdBridge(id)
	if etcdBridge != nil {
		s.EtcdBridgeMap.Delete(id)
		etcdBridge.Close()
		return true
	} else {
		return false
	}
}

// 根据id设置发送器
func (s *EtcdBridgeGroup) SetEtcdBridge(etcdBridge *EtcdBridge) {
	s.EtcdBridgeMap.Store(etcdBridge.ID, etcdBridge)
}

// 根据地址获取发送器
func (s *EtcdBridgeGroup) GetEtcdBridge(id int) (etcdBridge *EtcdBridge) {
	val, have := s.EtcdBridgeMap.Load(id)
	if have {
		etcdBridge = val.(*EtcdBridge)
	} else {
		etcdBridge = nil
	}
	return
}

// 用于控制每个发送器的 grpc 连接
type EtcdBridge struct {
	ID       int              //id
	Address  string           //连接地址
	Conn     *grpc.ClientConn //连接实体
	UserName string
	Password string
	IsSSL    bool
	Token    string
}

func (s *EtcdBridge) Init(id int, address string, userName, password string, isSSL bool) (err error) {
	s.ID = id
	s.Address = address
	s.UserName = userName
	s.Password = password
	s.IsSSL = isSSL
	var opts []grpc.DialOption
	ctx, _ := context.WithTimeout(context.Background(), 35*time.Second)
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
	s.Conn, err = grpc.DialContext(ctx, address, opts...)
	return
}
func (s *EtcdBridge) ReConnect() (err error) {
	if s.Conn != nil {
		s.Close()
	}
	s.Conn, err = grpc.Dial(s.Address, grpc.WithInsecure(), grpc.WithBlock())
	return
}
func (s *EtcdBridge) Close() (err error) {
	s.Conn.Close()
	s.Conn = nil
	return
}
