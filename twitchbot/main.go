package twitchbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/deanveloper/notezart"
	twitch "github.com/gempir/go-twitch-irc"
)

var client *twitch.Client
var config notezart.Config

// Initialize initializes the twitch bot
func Initialize() {

	initConfig()
	initMessages()

	client = twitch.NewClient(config.Twitch.Username, config.Twitch.Password)
	client.OnNewMessage(onMessage)
	client.OnConnect(onConnect)
	err := client.Connect()
	if err != nil {
		fmt.Println("Error in twitch connection:", err)
		os.Exit(1)
	}
}

// initializes global config variable or calls
// os.Exit(1) if an error occurs
func initConfig() {
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
}

func onConnect() {
	client.Join(config.Twitch.Username)
	client.Say(config.Twitch.Username, "Connected!")
}

func onMessage(channel string, user twitch.User, msg twitch.Message) {
	if msg.Text[0] != '!' {
		return
	}
	split := strings.SplitN(msg.Text, " ", 2)
	if len(split) == 1 {
		split = append(split, "")
	}
	cmd, args := split[0], split[1]

	// basically just an echo command for now
	switch cmd {
	case "!sr":
		requestCmd(user, channel, args)
	case "!songlist":
		listCmd(user, channel)
	}
}
