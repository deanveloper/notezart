package twitchbot

import (
	"fmt"

	"github.com/deanveloper/notezart/api"
	"github.com/deanveloper/notezart/messages"
	twitch "github.com/gempir/go-twitch-irc"
)

func requestCmd(user twitch.User, channel string, song string) error {
	vid, err := api.SearchForSong(config, song)
	vid.Requester = user.DisplayName
	if err == api.ErrNoVideoFound {
		client.Say(channel, messages.Message("SongNotFound", messages.MessageInput{
			User:  user,
			Video: api.Video{Title: song},
		}))
		return nil
	} else if err != nil {
		client.Say(channel, messages.Message("ErrorOccurred", messages.MessageInput{User: user}))
		return err
	}

	queue := api.SongQueue(channel)
	queue.Enqueue(vid)

	client.Say(channel, messages.Message("SongQueued", messages.MessageInput{
		User:  user,
		Video: vid,
		Songs: api.SongQueue(channel),
	}))

	return nil
}

func listCmd(user twitch.User, channel string) error {
	queue := api.SongQueue(channel)
	fmt.Println(queue)
	client.Say(channel, messages.Message("SongList", messages.MessageInput{
		User:  user,
		Songs: queue,
	}))
	return nil
}
