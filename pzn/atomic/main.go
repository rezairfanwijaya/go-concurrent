package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic digunakan untuk menghindari race condition atau sebagai pengganti dari mutex
// disini saya tidak akan menggunakan mutex sama sekali
func main() {
	var numberWithAtomic int32
	numberWithoutAtomic := 0

	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for i := 1; i <= 10; i++ {
				numberWithoutAtomic++ // race condition
			}
			group.Done()
		}()
	}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for i := 1; i <= 10; i++ {
				atomic.AddInt32(&numberWithAtomic, 1) // no race condition
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Printf("without atomic : %v\n", numberWithoutAtomic)
	fmt.Printf("with atomic : %v\n", numberWithAtomic)
}
