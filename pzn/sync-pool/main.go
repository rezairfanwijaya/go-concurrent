package main

import (
	"fmt"
	"sync"
	"time"
)

// pool akan menyimpan data
// dan kita bisa menggunakan data itu berkali kali

func main() {
	// deklarasi pool
	// pool dapat kita kasih nilai default
	// dengan menggunakan atribut new pada struct pool
	pool := sync.Pool{
		New: func() any {
			return "This default value"
		},
	}
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
			// if data != nil {
			// 	fmt.Printf("data : %v on iteration : %v\n", data, i)
			// }

			fmt.Printf("data : %v on iteration : %v\n", data, i)
			time.Sleep(1 * time.Second) // kasih jeda agar default value dapat dicetak
			pool.Put(data)
		}(i)
	}

	wg.Wait()
}
