package main

import (
	"log"
	"strconv"
)

func main() {
	// range channel dipakai ketika kita tidak tahu
	// berapa jumlah data yang akan dikirim ke channel
	// sehingga kita juga tidak tahu berapa kali kita
	// harus consume data tersebut
	// range channel akan berhenti ketika channel sudah di close
	// jadi jangan lupa untuk close channel

	word := make(chan string)

	go func() {
		defer close(word) // jangan lupa close channel, agar tidak deadlock
		for i := 1; i < 10; i++ {
			word <- "perulangan ke " + strconv.Itoa(i)
		}
	}()

	for value := range word {
		log.Println("menerima data : ", value)
	}

}
