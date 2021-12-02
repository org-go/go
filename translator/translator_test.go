package translator

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	f := Config{
		Proxy:       "http://127.0.0.1:80",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	s := New(f)
	result, err := s.Translate(`content`, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
