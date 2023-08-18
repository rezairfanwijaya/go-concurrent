package main

import (
	"log"
	"sync"
)

// in -> berarti channel hanya dapat menerima data
// syntax : chan<- tipe_data

// out -> berarti channel hanya dapat mengirim data
// syntax : <-chan tipe_data

var wg sync.WaitGroup

func main() {
	channel := make(chan int)
	defer close(channel)

	wg.Add(2)
	go DataIn(channel, 10)
	go DataOut(channel)
	wg.Wait()
}

func DataIn(channel chan<- int, data int) {
	channel <- data
	wg.Done()
}

func DataOut(channel <-chan int) {
	data := <-channel
	log.Println(data)
	wg.Done()
}
