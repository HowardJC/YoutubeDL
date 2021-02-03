package main

import (
	"encoding/json"
	"github.com/knadh/go-get-youtube/youtube"
	"net/url"
	"sync"
	"testing"
)

func TestYoutube(t *testing.T) {
	video, _ := youtube.Get("P-uhgIzHYYo")
	// download the video and write to file
	option := &youtube.Option{
		Rename: true, // rename file using video title
		Resume: true, // resume cancelled download
		Mp3:    true, // extract audio to MP3
	}
	video.Download(0, "video.mp4", option)

}

func TestDownload(t *testing.T) {
	v := VideoContext{}
	c := sync.Mutex{}
	finish := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			err := v.Download("music", c)
			if err != nil {
				t.Error(err)
			}
			finish <- true
		}()
		<-finish
	}
}

//TODO: remove sync mutex and just put it into the struct
func TestVideo(t *testing.T) {
	v := VideoContext{}
	c := sync.Mutex{}
	v.Download("music", c)

	parsed, err := url.ParseQuery(v.urlMeta)
	if err != nil {
		t.Error(err)
	}
	status := parsed.Get("status")
	if status != "ok" {
		t.Error("status does not exist?")
	}

	player := parsed.Get("player_response")
	if player == "" {
		t.Error("Does not exist")
	}
	var Clown responseInfo
	println(player)
	json.Unmarshal([]byte(player), &Clown)

	println(Clown.DashManifestURL)
}
