package helper

import (
	"fmt"
	"time"
)

func RandomID() string {
	randomID := fmt.Sprint(time.Now().UnixNano())
	return randomID
}

// func URIEncode(s string) string {
// 	return url.QueryEscape(s)
// }
