package main

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ServerStart    string `json:"serverStart"`
	ServerShutdown string `json:"serverShutdown"`
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("server_config.json", &configuration)
	fmt.Println(configuration)
	return configuration
}
