package main

import (
	"bookkeeping-server/cmd"
	"bookkeeping-server/config"
)

func main() {
	config.LoadConfigYaml()
	cmd.Run()
}
