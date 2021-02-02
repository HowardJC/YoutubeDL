package YoutubeDownloader

import "net/http"

func (video *Video) Download(rep string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
}
