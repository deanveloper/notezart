package main

import (
	"github.com/deanveloper/notezart"
	twitch "github.com/gempir/go-twitch-irc"
)

func requestCmd(user twitch.User, channel string, song string) error {
	vid, err := notezart.SearchForSong(config, song)
	if err == notezart.ErrNoVideoFound {
		client.Say(channel, message("SongNotFound", MessageInput{User: user}))
		return nil
	} else if err != nil {
		client.Say(channel, message("ErrorOccurred", MessageInput{User: user}))
		return err
	}

	client.Say(channel, message("SongQueued", MessageInput{User: user, Video: vid}))

	return nil
}
