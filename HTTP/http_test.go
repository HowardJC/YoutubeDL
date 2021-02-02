package HTTP

import (
	"testing"
)

func TestRequest(t *testing.T) {
	body, _ := GetRequest("https://youtu.be/RpQWQ7FCjLI?list=RDMM")
	t.Logf(string(body))
}
