package main

import (
	"os"

	configuration "./configuration"
	server "./server"
)

var (
	configPath string
)

func main() {

	config := configuration.ParseConfig(os.Getenv("CONFIG_PATH"))
	server.StartServer(config)

}
