package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	twitch "github.com/gempir/go-twitch-irc"
)

var client *twitch.Client

var config struct {
	Username string `json:"username"`
	OAuth    string `json:"oauth"`
}

func main() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error while reading config.json:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error while parsing config.json:", err)
		os.Exit(1)
	}
	client = twitch.NewClient(config.Username, config.OAuth)
	client.OnNewMessage(onMessage)
	client.OnConnect(onConnect)
	client.Join(config.Username)
	err = client.Connect()
	if err != nil {
		fmt.Println("Error in connection:", err)
		os.Exit(1)
	}
}

func onConnect() {
	client.Say(config.Username, "Connected!")
}

func onMessage(channel string, user twitch.User, msg twitch.Message) {
	if msg.Text[0] != '!' {
		return
	}
	split := strings.SplitN(msg.Text, " ", 2)
	cmd, args := split[0], split[1]

	// basically just an echo command for now
	switch cmd {
	case "!sr":
		client.Say(channel, user.DisplayName+" has requested "+args)
	}
}
