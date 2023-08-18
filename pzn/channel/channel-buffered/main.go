package main

import "log"

func main() {
	// jumlah buffer menentukan seberapa banyak channel itu bisa menampung data
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "hallo1"
	channel <- "hallo2"
	channel <- "hallo3"
	log.Println(<-channel)

	log.Println(cap(channel))
	log.Println(len(channel))

	log.Println("selesai")
}
