package twitchbot

import (
	"bytes"
	"fmt"
	"os"
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
	Songs *notezart.VideoList
}

var messages *template.Template

// initializes global messages variable or calls
// os.Exit(1) if an error occurs
func initMessages() {
	tmpl, err := template.New("").ParseFiles("defaultMessages.tmpl")
	if err != nil {
		fmt.Println("Error while parsing defaultMessages.tmpl:", err)
		os.Exit(1)
	}
	messages = tmpl
}

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
