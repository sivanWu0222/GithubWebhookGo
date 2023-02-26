package conf

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func init() {
	config = viper.New()
	config.AddConfigPath("./conf")
	config.SetConfigName("config")
	config.SetConfigType("ini")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("找不到配置文件")
		} else {
			log.Fatalln(err.Error())
		}
	}
}

func GetStringContent(key string) string {
	return config.GetString(key)
}

func GetIntContent(key string) int {
	return config.GetInt(key)
}
