package twitchbot

import (
	"fmt"

	"github.com/deanveloper/notezart"
	twitch "github.com/gempir/go-twitch-irc"
)

func requestCmd(user twitch.User, channel string, song string) error {
	vid, err := notezart.SearchForSong(config, song)
	vid.Requester = user.DisplayName
	if err == notezart.ErrNoVideoFound {
		client.Say(channel, message("SongNotFound", MessageInput{User: user}))
		return nil
	} else if err != nil {
		client.Say(channel, message("ErrorOccurred", MessageInput{User: user}))
		return err
	}

	queue := notezart.SongQueue(channel)
	queue.Enqueue(vid)

	client.Say(channel, message("SongQueued", MessageInput{
		User:  user,
		Video: vid,
		Songs: notezart.SongQueue(channel),
	}))

	return nil
}

func listCmd(user twitch.User, channel string) error {
	queue := notezart.SongQueue(channel)
	fmt.Println(queue)
	client.Say(channel, message("SongList", MessageInput{
		User:  user,
		Songs: queue,
	}))
	return nil
}
