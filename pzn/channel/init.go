package main

import "fmt"

func main() {
	// init channel
	chanNumber := make(chan int)

	go func() {
		// masukan data ke channel
		chanNumber <- 90
	}()

	// mengambil data dari channel
	res := <-chanNumber
	fmt.Println(res)

	// channel sebagai parameter
	// fmt.Println(<-chanNumber)

}
