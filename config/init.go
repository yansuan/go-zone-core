package config

import (
	"github.com/spf13/viper"
)

// 设置默认读取目录
func Init(dirs ...string) {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("etc/")
	//
	for _, d := range dirs {
		viper.AddConfigPath(d)
	}
}
