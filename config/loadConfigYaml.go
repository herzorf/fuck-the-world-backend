package config

import (
	"fuck-the-world/unit"
	"github.com/spf13/viper"
	"os"
)

func LoadConfigYaml() {
	env := os.Getenv("APP_ENV") // APP_ENV 应该设置为 "development" 或 "production"
	if env == "" {
		env = "development" // 默认环境
	}
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		unit.HandleError("viper 读取错误", err)
	}
}
