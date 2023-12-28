package sample

import (
	"fmt"

	"time"

	"github.com/thuongtruong1009/gouse/cache"
	"github.com/thuongtruong1009/gouse/date"
)

func cacheLocal() {
	newCache := cache.NewLocal()
	newCache.SetLocal("key1", "local cache value 1")
	newCache.SetLocal("key2", "local cache value 2")
	newCache.SetLocal("key3", "local cache value 3")

	all := newCache.AllLocal()
	println("All local cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := newCache.GetLocal("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Getted key 1:", getted1)

	newCache.DelLocal("key2")

	all = newCache.AllLocal()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	newCache.FlushLocal()

	all = newCache.AllLocal()
	fmt.Println("All local cache values (after flush):", all)
}

func cacheTmp() {
	newCache := cache.NewTmp(date.ToSecond(3))
	newCache.SetTmp("key1", "temp cache value 1", date.ToSecond(3))
	newCache.SetTmp("key2", "temp cache value 2", date.ToSecond(6))
	newCache.SetTmp("key3", "temp cache value 3", date.ToSecond(6))

	all := newCache.AllTmp()
	println("All temp cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	getted := newCache.GetTmp("key1")
	fmt.Println("Getted key 1 (before expires):", getted)

	time.Sleep(date.ToSecond(4))

	getted = newCache.GetTmp("key1")
	fmt.Println("Getted key 1 (after expires):", getted)

	newCache.DelTmp("key2")

	all = newCache.AllTmp()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	newCache.FlushTmp()

	all = newCache.AllTmp()
	fmt.Println("All temp cache values (after flush):", all)
}
