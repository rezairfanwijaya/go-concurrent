package main

import (
	"fmt"
	"sync"
)

// sync map adalah tipe data map yang dikhususkan untuk digunakan pada goroutine
// karena aman dari race condition

func main() {
	mapData := &sync.Map{}
	waitGroup := &sync.WaitGroup{}

	var addToMap = func(value int) {
		mapData.Store(value, value)
	}

	// saya akan masukan 500 data ke map
	for i := 1; i <= 500; i++ {
		waitGroup.Add(1)

		go func(i int) {
			defer waitGroup.Done()
			addToMap(i)
		}(i)
	}

	waitGroup.Wait()

	// saya akan mengeluarkan data dari map
	mapData.Range(func(key, value any) bool {
		if value.(int)%2 == 0 {
			fmt.Printf("key : %v  |  value : %v\n", key, value)
		}
		return true
	})

}
