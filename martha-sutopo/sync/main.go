package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go CetakNama("Hi")
	go CetakNama("Hallo")
	wg.Wait()

	fmt.Println("program selesai")
}

func CetakNama(s string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s ke %d\n", s, i)
		<-time.After(time.Second)
	}
}
