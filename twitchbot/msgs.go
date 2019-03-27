package main

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"text/template"

	"github.com/deanveloper/notezart"

	twitch "github.com/gempir/go-twitch-irc"
)

// MessageInput represents the input that is passed
// into the messages
type MessageInput struct {
	User  twitch.User
	Video notezart.Video
}

var messages *template.Template

// Message returns a message from the defaultMessages.txt with
// the given template inputs.
func message(key string, inputs MessageInput) string {
	var buf bytes.Buffer
	err := messages.ExecuteTemplate(&buf, key, inputs)
	if err != nil {
		fmt.Println("template error:", err)
		debug.PrintStack()
	}
	return string(buf.Bytes())
}
