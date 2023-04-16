package twitchbot

import (
	"fmt"
	"os"
	"strings"

	"github.com/deanveloper/notezart/api"
	twitch "github.com/gempir/go-twitch-irc"
)

var client *twitch.Client
var config api.Config

// Start initializes the twitch client and registers needed events.
func Start(cfg api.Config) {
	config = cfg
	client = twitch.NewClient(config.Twitch.Username, config.Twitch.Password)
	client.OnNewMessage(onMessage)
	client.OnConnect(onConnect)
	err := client.Connect()
	if err != nil {
		fmt.Println("Error in twitch connection:", err)
		os.Exit(1)
	}
}

// Join joins a channel
func Join(user string) {
	client.Join(user)
}

func onConnect() {
	// always connect to own channel
	Join(config.Twitch.Username)
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
