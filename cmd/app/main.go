package main

import (
	"fmt"

	"github.com/Hulla-Hoop/LRU_Cache/internal/cache/lrucache"
)

func main() {
	c := lrucache.New(3)
	f := c.Add("key1", "value1")
	f2 := c.Add("key1", "value1")
	fmt.Println(f, f2)
}
