package sample

import (
	"fmt"

	"time"

	"github.com/thuongtruong1009/gouse/config/cache"
	"github.com/thuongtruong1009/gouse/date"
)

func confCacheTmp() {
	cache.SetTmp("key", "temp cache value")

	getted, err := cache.GetTmp("key")
	if err != nil {
		panic(err)
	}

	fmt.Println(getted)
}

func confCacheLocal() {
	cache.SetLocal("key", "local in-memory cache value", date.ToSecond(3))

	getted := cache.GetLocal("key")
	fmt.Println(getted)

	time.Sleep(date.ToSecond(4))

	getted = cache.GetLocal("key")
	fmt.Println(getted)
}
