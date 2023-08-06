package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	chanName := make(chan string)

	wg.Add(1)
	go CetakNama("Hallo", chanName)
	res := <-chanName
	wg.Wait()
	fmt.Println(res)
}

func CetakNama(s string, chanName chan string) {
	defer wg.Done()
	chanName <- s
}
