package config

import (
	"bookkeeping-server/unit"
	"github.com/spf13/viper"
)

func LoadConfigYaml() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/DESKTOP/bookkeeping-server/")
	err := viper.ReadInConfig()
	if err != nil {
		unit.HandleError("viper 读取错误", err)
	}
}
