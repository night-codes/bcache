package main

import (
	"fmt"
	"github.com/mirrr/bcache"
	"time"
)

func main() {
	cache := bcache.Create()
	cache.Updater(func(key string) interface{} {
		if key == "mykey" {
			return 777
		}
		return nil
	})
	cache.Set("mykey2", 888)

	fmt.Println(cache.Get("mykey"))  // 777
	fmt.Println(cache.Get("mykey2")) // 888

	time.Sleep(time.Second * 2)

	fmt.Println(cache.Get("mykey"))  //  777
	fmt.Println(cache.Get("mykey2")) //  <nil>
}
