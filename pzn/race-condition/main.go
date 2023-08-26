package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// mutex
	Problem()
	Solution()

	// RWMutex
	bank := BankBallance{}
	// saya akan siapkan 100 goroutine
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// tiap goroutine akan menaikan balance di bank
			// sebanyak 100
			for i := 1; i <= 100; i++ {
				bank.AddBallance(1)
				log.Println("latest add : ", bank.GetBallance())
			}
		}()
	}
	wg.Wait()

	// total balance harusny ada 100 * 100 = 10.000
	log.Println("latest balance in bank : ", bank.GetBallance())
}
