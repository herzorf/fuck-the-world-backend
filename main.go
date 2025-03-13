package main

import (
	"bookkeeping-server/cmd"
	"bookkeeping-server/unit"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()
	unit.HandleError("viper 读取错误", err) // Find and read the config file
	cmd.Run()
}
