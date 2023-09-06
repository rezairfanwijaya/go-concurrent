package main

import (
	"fmt"
	"sync"
)

// pool akan menyimpan data
// dan kita bisa menggunakan data itu berkali kali

func main() {
	pool := sync.Pool{}
	var wg sync.WaitGroup

	// put data to pool
	pool.Put("one")
	pool.Put("two")
	pool.Put("three")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data := pool.Get()
			if data != nil {
				fmt.Printf("data : %v on iteration : %v\n", data, i)
			}
			pool.Put(data)
		}(i)
	}

	wg.Wait()
}
