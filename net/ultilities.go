package net

import (
	"net/http"
	"net/url"
	"time"

	"github.com/thuongtruong1009/gouse/console"
)

func Open(url string) {
	console.Cmd("explorer " + url)
}

func Encode(s string) string {
	return url.QueryEscape(s)
}

func Decode(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}

func Check(url string) (bool, error) {
	_, err := http.Get(url)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckWithStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func Header(url string) (http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

func ConnectTime(url string) (float64, error) {
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elapsedTime := time.Since(startTime)
	return elapsedTime.Seconds(), nil
}
