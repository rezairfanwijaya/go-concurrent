package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100000; i++ {
		group.Add(1)
		go helloWorld(group, i)
	}

	group.Wait()
	fmt.Println("Done")

}

func helloWorld(group *sync.WaitGroup, iteration int) {
	defer group.Done()
	fmt.Println("Hello World : ", iteration)

	// latency
	<-time.After(time.Second * 1)
}
