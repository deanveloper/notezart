package api

import (
	"fmt"
	"html"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var songLists = make(map[string]*VideoList)

// SongQueue returns the given channel's video list.
// If the channel does not yet have a video list, it is created.
func SongQueue(channel string) *VideoList {
	list, ok := songLists[channel]
	if !ok {
		list = new(VideoList)
		songLists[channel] = list
	}
	return list
}

// SearchForSong searches for a song on YouTube. `query` should be a
// standard youtube search. Returns the top video.
func SearchForSong(config Config, query string) (Video, error) {
	httpClient := &http.Client{
		Transport: &transport.APIKey{Key: config.Youtube.Key},
	}
	service, err := youtube.New(httpClient)
	if err != nil {
		return Video{}, fmt.Errorf("error initializing service in requestCmd: %v", err)
	}

	call := service.Search.List("id,snippet").Type("video").Q(query).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return Video{}, fmt.Errorf("error searching for %q in requestCmd: %v", query, err)
	}

	if len(response.Items) == 0 {
		return Video{}, ErrNoVideoFound
	}

	result := response.Items[0]
	return Video{
		ID:     result.Id.VideoId,
		Title:  html.UnescapeString(result.Snippet.Title),
		Artist: html.UnescapeString(result.Snippet.ChannelTitle),
	}, nil
}
