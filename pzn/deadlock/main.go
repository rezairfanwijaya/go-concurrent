package main

import (
	"fmt"
	"sync"
	"time"
)

// deadlock terjadi ketika ada 2 goroutine atau lebih
// akan menggunakan data untuk di lock dari masing-masing goroutine
// tapi data tersebut masih dilock oleh masing-masing goroutine
// contoh
/**
goroutine a :
1. lock data Z
2. lock data Y

goroutine b :
1. lock data Y
2. lock data z

goroutine a butuh data Y dari goroutine b
goroutine b butuh data  z dari goroutine a

dan akan terjadi hang atau deadlock
*/

// study case
type BankBRIUser struct {
	sync.Mutex
	Name    string
	Balance int
}

// fungsi lock
func (b *BankBRIUser) LockUser() {
	b.Mutex.Lock()
}

// fungsi unlock
func (b *BankBRIUser) UnlockUser() {
	b.Mutex.Unlock()
}

// fungsi ubah balance
func (b *BankBRIUser) ChangeBalance(amount int) {
	b.Balance += amount
}

// fungsi transfer
func (b *BankBRIUser) Transfer(user1 *BankBRIUser, user2 *BankBRIUser, amount int) {
	defer wg.Done()
	// lock user 1
	user1.LockUser()
	fmt.Println("lock user 1 : ", user1.Name)
	user1.ChangeBalance(-amount)
	// user1.UnlockUser() // uncomment agar tidak deadlock

	<-time.After(1 * time.Second)

	// lock user 2
	user2.LockUser()
	fmt.Println("lock user 2 : ", user2.Name)
	user2.ChangeBalance(amount)
	// user2.UnlockUser() // uncomment agar tidak deadlock

	// kode ini akan membuat deadlock
	user1.Unlock()
	user2.Unlock()
}

var wg sync.WaitGroup

func main() {
	reza := BankBRIUser{
		Name:    "reza",
		Balance: 100,
	}

	irfan := BankBRIUser{
		Name:    "irfan",
		Balance: 100,
	}

	wg.Add(2)
	// reza dan irfan akan sama sama transfer
	go reza.Transfer(&reza, &irfan, 90)
	go irfan.Transfer(&irfan, &reza, 10)
	wg.Wait()

	fmt.Printf("user 1 %v balance %v\n", reza.Name, reza.Balance)
	fmt.Printf("user 2 %v balance %v\n", irfan.Name, irfan.Balance)
}
