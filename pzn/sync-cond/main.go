package main

import (
	"fmt"
	"sync"
	"time"
)

// sync cond digunakan untuk membuat seperti antrian
// pada goroutine ketika akan mengkonsumi data

func main() {
	group := &sync.WaitGroup{}
	mutex := sync.Mutex{}
	condition := sync.NewCond(&mutex)
	isReady := false
	buffer := 7
	data := make(chan int, 11)

	for i := 1; i < buffer; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			time.Sleep(2 * time.Second)

			condition.L.Lock()

			data <- i
			isReady = true
			fmt.Printf("data %v is ready!\n", i)
			condition.Signal() // untuk memberitahu bahwa data sudah siap

			condition.L.Unlock()
		}(i)
	}

	for i := 1; i < buffer; i++ {
		group.Add(1)
		go func(data chan int) {
			defer group.Done()

			condition.L.Lock()

			for !isReady {
				fmt.Println("sedang menunggu data")
				condition.Wait() // untuk memberitahu goroutine suruh menunggu
			}

			fmt.Printf("data %v diterima\n", <-data)

			condition.L.Unlock()
		}(data)
	}

	group.Wait()
}
