package baseconfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"time"
)

var CONFIG Config

func init() {
	currentDir, err := os.Getwd()
	base := filepath.Base(currentDir)
	if base == "test" {
		currentDir += "/.."
	}
	// 设置配置文件名和路径
	viper.AddConfigPath(currentDir)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// 自动从环境变量中读取配置
	viper.AutomaticEnv()
	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("无法读取配置文件:", err)
		return
	}
	// 获取配置项的值
	CONFIG.Host = viper.GetString("SERVER_HOST")
	CONFIG.Debug = viper.GetBool("SERVER_DEBUG")
	CONFIG.ServerPublicKey = viper.GetString("SERVER_SERVERPUBLICKEY")
	CONFIG.ServerPrivateKey = viper.GetString("SERVER_SERVERPRIVATEKEY")
	CONFIG.DBPath = viper.GetString("DB_PATH")
	CONFIG.PageSize = viper.GetInt("PAGESIZE")

}

type Config struct {
	Host                 string //监听地址
	Debug                bool
	ServerPublicKey      string
	ServerPrivateKey     string
	TokenDuration        time.Duration
	DBPath               string
	OpenSSL              bool
	PageSize             int
	PeerNowEndPoint      string
	PeerNowKey           string
	PeerEndPoints        string
	PeerPublicKey        string
	PeerPrivateKey       string
	PeerCA               string
	PeerOpenSSL          bool
	PeerClientAuth       bool
	PeerClientPublicKey  string
	PeerClientPrivateKey string
}
