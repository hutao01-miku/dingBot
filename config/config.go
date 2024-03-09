// config_loader.go

package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ClientID      string
	ClientSecret  string
	APIURL        string
	APIKey        string
	SystemMessage string
)

func init() {
	// 设置配置文件的名字和类型
	viper.SetConfigName("config") // 不需要文件后缀
	viper.SetConfigType("yaml")

	// 添加配置文件的查找路径
	viper.AddConfigPath("config") // 相对路径，根据实际情况修改

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 使用 viper 读取配置
	ClientID = viper.GetString("clientId")
	ClientSecret = viper.GetString("clientSecret")
	APIURL = viper.GetString("apiURL")
	APIKey = viper.GetString("apiKey")
	SystemMessage = viper.GetString("systemMessage")

	// 在这里可以添加其他配置项...
}
