package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// buffer ini akan membatasi channel untuk
	// menampung data
	// dengan jumlah maximal data adalah
	// buffer + 1
	buffer := 5
	chanNumbers := make(chan int, buffer)

	wg.Add(1)
	// menerima
	go func() {
		for chanNumber := range chanNumbers {
			chanNumber = <-chanNumbers
			fmt.Println("terima data ============ ", chanNumber)
		}
	}()

	// mengirim
	for i := 1; i <= 10; i++ {
		fmt.Println("kirim data ", i)
		chanNumbers <- i
	}
	wg.Done()
	wg.Wait()
}
