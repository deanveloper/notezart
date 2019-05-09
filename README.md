
:fire_engine: STILL INDEV :fire_engine:

# Notezart

A better music bot for better streamers.

Will have the following features:
 - [x] Request songs via chat
 - [ ] Request songs via web interface
 - [ ] Web interface for viewing the song list
 - [ ] Web interface for listening to the queue
 - [ ] Blacklist certain songs, artists, or keywords
 - [ ] Easy-to-use API to allow developers to make their own forks

Since I need to pay for hosting, I'll have a premium plan
which will allow streamers to have configurable responses.

Will probably only allow songs from YouTube. I'm not 100% sure on that yet, though.

### Directory Structure

```
notezart/: main package, connects everything together
	api/: 		backend and base API
	messages/:  handles templating for messages and webpages
	twitchbot/: twitch chat frontend
	web/:       [coming soon] web server
```
