package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// select dipakai ketika kita akan mengambil data dari 2 channel yang berbeda
	// didalam satu loop
	// kalau memakai for range atau range channel kita hanya terbatas
	// pada 1 channel saja, tidak bisa multiple channel

	channel1 := make(chan string)
	channel2 := make(chan string)

	wg.Add(2)
	go func() {
		defer close(channel1)
		defer wg.Done()
		for i := 1; i < 11; i++ {
			// set delay untuk masuk ke default select
			if i == 5 || i == 9 {
				<-time.After(3 * time.Millisecond)
			}

			data := fmt.Sprintf("Data ke %d dari channel 1", i)
			channel1 <- data
		}
	}()

	go func() {
		defer close(channel2)
		defer wg.Done()
		for i := 1; i < 11; i++ {
			// set delay untuk masuk ke default select
			if i == 2 || i == 10 {
				<-time.After(3 * time.Millisecond)
			}

			data := fmt.Sprintf("Data ke %d dari channel 2", i)
			channel2 <- data
		}
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			log.Println(data)
			counter++
		case data := <-channel2:
			log.Println(data)
			counter++
		default:
			log.Println("menunggu data")
		}

		if counter == 20 {
			break
		}
	}

	wg.Wait()
}
