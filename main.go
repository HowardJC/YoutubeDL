package main

import (
	"YoutubeDownloader/HTTP"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type VideoContext struct {
	error   error
	data    []byte
	url     string
	urlMeta string
}

func (v *VideoContext) Download(Folder string, mu sync.Mutex) error {
	//TODO: make seperate struct/function to create the directorylist and do cleaning for the names also take the 2 minutes to choose a url and not have a dummy
	body, err := HTTP.GetRequest("https://www.youtube.com/watch?v=C0DPdy98e4c")
	if err != nil {
		return nil
	}

	if string(body) == "" {
		return nil
	} else if Folder != "" {
		os.Mkdir(Folder, 0777)
	}

	mu.Lock()
	DirList, _ := ioutil.ReadDir(Folder)
	//TODO: use regex and format string
	filename := fmt.Sprintf("%s/File_%d", Folder, len(DirList))
	mu.Unlock()
	file, err := os.Create(filename)
	_, err = file.Write(body)
	v.data = body
	//TODO: remove dummy after ii finish figuring out what is returned from this
	v.url = fmt.Sprintf("https://youtube.com/get_video_info?video_id=%s&eurl=https://youtube.googleapis.com/v/%s", "3Xjdxi3ILxU", "3Xjdxi3ILxU")
	if err != nil {
		return errors.New("Error writings")
	}

	body, err = HTTP.GetRequest(v.url)

	v.urlMeta = string(body)
	return nil
}

func (v VideoContext) VideoMetadata() {
	println(v.data)

}
