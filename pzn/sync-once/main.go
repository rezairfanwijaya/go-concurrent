package main

import (
	"fmt"
	"sync"
)

var (
	CounterWithOnce    int
	CounterWithoutOnce int
)

func OnlyOnceExecution() {
	CounterWithOnce++
}

func ExecuteEveryTime() {
	CounterWithoutOnce++
}

func main() {
	// synce once hanya akan memperbolehkan function
	// dieksekusi sekali saja
	// ketika sudah dieksekusi maka function itu tidak akan dieksekusi lagi
	// alias sudah expired

	// disini saya akan peragakan
	// ketika variabel di ubah menggunakan once
	// dan tidak menggunakan once
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	once := &sync.Once{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// with once
			once.Do(OnlyOnceExecution)

			// without once
			mtx.Lock()
			ExecuteEveryTime()
			mtx.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter with once: ", CounterWithOnce)
	fmt.Println("counter without once: ", CounterWithoutOnce)

}
