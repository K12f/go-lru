package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	//lru( latest recently used) 算法

	lru := NewLRU(10)

	for i := 0; i < 10; i++ {
		lru.Set(strconv.Itoa(i), i)
	}

	fmt.Println(lru.Get("1"))

	var next *list.Element
	for e := lru.list.Front(); e != nil; e = next {
		next = e.Next()
		fmt.Println(e.Value)
	}

	for key, value := range lru.cache {
		fmt.Println(key, value.Value)
	}

	fmt.Println(lru.cache)

}
