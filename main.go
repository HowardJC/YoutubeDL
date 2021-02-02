package main

import (
	"YoutubeDownloader/HTTP"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func Download(Folder string, mu sync.Mutex) error {
	//TODO: make seperate struct/function to create the directorylist and do cleaning for the names also take the 2 minutes to choose a url and not have a dummy
	body, err := HTTP.GetRequest("https://www.youtube.com/watch?v=UATNrUebey0&list=RDMM&index=16")
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
	if err != nil {
		return errors.New("Error writing")
	}
	return nil

}
