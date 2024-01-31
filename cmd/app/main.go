package main

import (
	"fmt"
	"time"

	"github.com/Hulla-Hoop/LRU_Cache/internal/cache/lrucache"
)

func main() {

	c := lrucache.New(3)

	for i := 0; i < 5; i++ {
		go func(i int) {
			f1 := c.Add("key1", "value1")
			fmt.Println("горутина---", i, "--", f1)
			f2 := c.Add("key2", "value1")
			fmt.Println("горутина---", i, "--", f2)
			data, d := c.Get("key1")
			fmt.Println("горутина---", i, "--", data, d)
			data2, d2 := c.Get("key2")
			fmt.Println("горутина---", i, "--", data2, d2)

		}(i)
	}

	time.Sleep(10 * time.Second)
}
