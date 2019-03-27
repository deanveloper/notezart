package notezart

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// ErrNoVideoFound represents when a video which was searched for cannot be found
var ErrNoVideoFound = errors.New("no video found")

// Video is a struct which represents a Youtube Video
type Video struct {
	ID     string
	Title  string
	Artist string
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
	return Video{ID: result.Id.VideoId}, nil
}
