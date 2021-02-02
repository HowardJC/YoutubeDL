package main

import (
	"sync"
	"testing"
)

func TestDownload(t *testing.T) {
	c := sync.Mutex{}
	finish := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			err := Download("music", c)
			if err != nil {
				t.Error(err)
			}
			finish <- true
		}()
		<-finish
	}
}
