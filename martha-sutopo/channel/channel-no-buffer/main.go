package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chanNumber := make(chan int)

	wg.Add(2)
	go func(i int) {
		defer wg.Done()
		chanNumber <- i
	}(10)

	var funcChanNumber = func(i int) {
		defer wg.Done()
		chanNumber <- i
	}
	go funcChanNumber(90)

	res := <-chanNumber
	final := <-chanNumber
	wg.Wait()

	fmt.Println(res)
	fmt.Println(final)
}
