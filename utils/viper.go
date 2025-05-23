package utils

import "github.com/spf13/viper"

func InitViper() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic("配置文件加载失败" + err.Error())
	}
}
