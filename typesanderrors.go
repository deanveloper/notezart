package notezart

import (
	"errors"
)

var (
	// ErrNoVideoFound represents when a video which was searched for cannot be found
	ErrNoVideoFound = errors.New("no video found")

	// ErrEmptyList represents when one tries to access
	// an element from an empty VideoList
	ErrEmptyList = errors.New("list is empty")
)

// Config is the Go representation of the config.json file
type Config struct {
	Twitch struct {
		Username string
		Password string
	}
	Youtube struct {
		Key string
	}
}

// Video is a struct which represents a Youtube Video
// requested by a user
type Video struct {
	ID        string
	Title     string
	Artist    string
	Requester string
}

// VideoList represents an object that is used to interact
// with a channel's list of queued videos
type VideoList struct {
	list []Video
}

// Len returns the length of the list
func (l *VideoList) Len() int {
	return len(l.list)
}

// Enqueue adds a video to the end of the queue
func (l *VideoList) Enqueue(v Video) {
	if l.list == nil {
		l.list = make([]Video, 0, 1)
	}
	l.list = append(l.list, v)
}

// Slice returns the list of videos as a slice. If indices
// are out-of-bounds, they are coerced to be in-bounds. This
// means that the length of the returned slice may be less
// than (end - start).
func (l *VideoList) Slice(start, end int) []Video {
	if start < 0 {
		start = 0
	}
	if start > l.Len() {
		start = l.Len()
	}
	if end < 0 {
		end = 0
	}
	if end > l.Len() {
		end = l.Len()
	}
	if start >= end {
		return nil
	}
	newSlice := make([]Video, end-start)
	copy(newSlice, l.list[start:end])
	return newSlice
}

// Dequeue removes a video from the beginning of the queue,
// or an error if there is no video to remove
func (l *VideoList) Dequeue() (Video, error) {
	if l.Len() == 0 {
		return Video{}, ErrEmptyList
	}

	vid := l.list[0]
	l.list = l.list[1:]

	return vid, nil
}

// Clear removes all songs from the list
func (l *VideoList) Clear() {
	l.list = nil
}
