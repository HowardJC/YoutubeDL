package main

type responseInfo struct {
	streamingData struct {
		expiresInSeconds string `json:"responseContext"`
	}
	cards struct {
		trackingParams string
	}
}
