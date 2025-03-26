package main

import (
	"fuck-the-world/cmd"
	"fuck-the-world/config"
)

func main() {
	config.LoadConfigYaml()
	cmd.Run()
}
