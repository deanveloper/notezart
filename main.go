package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deanveloper/notezart/api"
	"github.com/deanveloper/notezart/twitchbot"
)

func main() {
	cfg := getConfig()

	twitchbot.Start(cfg)
}

func getConfig() api.Config {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error while reading config.json:", err)
		os.Exit(1)
	}

	var cfg api.Config

	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		fmt.Println("Error while parsing config.json:", err)
		os.Exit(1)
	}

	return cfg
}
