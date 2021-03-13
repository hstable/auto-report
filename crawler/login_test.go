package crawler

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func TestGetLt(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	var client = http.Client{
		Jar: jar,
	}
	lt, err := getLt(client)
	fmt.Println(lt, err)
}

func TestLogin(t *testing.T) {
	fmt.Println("")
}
