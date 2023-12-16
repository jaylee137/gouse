package helper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func RandomID() string {
	randomID := fmt.Sprint(time.Now().UnixNano())
	return randomID
}

func UUID() string {
	uuid := uuid.New()
	return uuid.String()
}

// func URIEncode(s string) string {
// 	return url.QueryEscape(s)
// }
