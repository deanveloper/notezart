package notezart

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
