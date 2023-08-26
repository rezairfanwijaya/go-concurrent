package main

import (
	"sync"
)

// mutex hanya akan mengunci aatu kegiatan saja
// entah itu untuk read ataupun write
// pada file solution.go mutex digunakan untuk
// mengunci kegiatan write atau ubah data
// lalu bagaimana jika ada 2 kegiatan
// yaitu read dan write?
// kita bisa menggunakan RWMutex

// study case
// BankBallance memiliki 2 fungsi
// yaitu AddBalance dan GetBallance
// saya akan buat struct BankBallance
type BankBallance struct {
	Ballance int
	RWM      sync.RWMutex // saya declare rwmutex disini
}

func (b *BankBallance) AddBallance(amount int) {
	b.RWM.Lock() // kunci untuk kegiatan ubah data
	b.Ballance += amount
	b.RWM.Unlock() // buka kunci untuk kegiatan ubah data
}

func (b *BankBallance) GetBallance() int {
	b.RWM.RLock() // kunci untuk kegiatan baca data
	ballance := b.Ballance
	b.RWM.RUnlock() // buka kunci untuk kegiatan baca data
	return ballance
}
