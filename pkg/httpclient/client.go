package httpclient

import (
	"net/http"
	"time"
)

var Client = &http.Client{
    Timeout: 30 * time.Second,
}

func GetUserAgent() string {
    return "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
}
