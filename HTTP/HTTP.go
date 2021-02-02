package HTTP

import (
	"io/ioutil"
	"net/http"
	"time"
)

func GetRequest(url string) (body []byte, err error) {
	client := http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(r.Body)
	return body, nil

}
