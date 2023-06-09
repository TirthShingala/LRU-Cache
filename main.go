package main

import (
	"fmt"

	"github.com/TirthShingala/LRU-Cache/lrucache"
)

func main() {
	lruCache := lrucache.Constructor(2)

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1)) // Output: 1

	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2)) // Output: -1

	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1)) // Output: -1
	fmt.Println(lruCache.Get(3)) // Output: 3
	fmt.Println(lruCache.Get(4)) // Output: 4
}
