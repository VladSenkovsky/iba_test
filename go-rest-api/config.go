package main

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ServerStart    string `json:"serverStart"`
	ServerShutdown string `json:"serverShutdown"`
}

func GetConfigTime() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("server_config.json", &configuration)
	return configuration
}
