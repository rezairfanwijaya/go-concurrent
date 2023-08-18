package main

import (
	"log"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	log.Println(<-channel)
}

func GiveMeResponse(channel chan string) {
	<-time.After(2 * time.Second)
	channel <- "hello this is address data"
}
