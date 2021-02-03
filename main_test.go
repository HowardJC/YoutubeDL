package main

import (
	"encoding/json"
	"fmt"
	"github.com/knadh/go-get-youtube/youtube"
	"github.com/valyala/fastjson"
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
	parsedtest, _ := url.Parse(v.urlMeta)
	println(parsedtest.Query().Get("status"))
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

	println("Hey")
	println(player)
	println(player)
	//if err:=json.Unmarshal([]byte(player), &Clown); err!=nil{
	//	t.Error("Unmarshal error?")
	//}
	//youtube json changes and too long?
	var p fastjson.Parser
	Parsed, _ := p.Parse(v.urlMeta)
	fmt.Println(Parsed)

}
