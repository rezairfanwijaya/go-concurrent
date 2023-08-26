package main

import (
	"log"
	"sync"
)

func Solution() {
	// untuk mengatasi race condition, kita harus membatasi akses
	// ke pada data yang akan diubah dengan aturan
	// bahwa hanya diperbolehkan satu goroutine saja yang
	// dapat mengakses data dalam satu waktu
	// ini dapat dilakukan dengan mutex

	// saya akan membuat variable yang dapat diakses oleh semua goroutine
	counter := 0

	// saya akan siapkan mutex untuk
	// digunakan dalam melakukan locking dan unlocking
	var mtx sync.Mutex

	// saya akan buat 2000 goroutine
	wg.Add(2000)
	for i := 1; i <= 2000; i++ {
		go func() {
			defer wg.Done()
			// lalu setiap go routine
			// akan menaikan counter sebanyak 500
			for j := 1; j <= 500; j++ {
				// saya akan lock, sehingga hanya satu
				// goroutine yang bisa mengubah data counter
				mtx.Lock()

				// lalu ubah value
				counter++

				// saya akan unlock
				// untuk membolehkan goroutine selanjutnya mengubah data
				// counter
				mtx.Unlock()
			}
		}()
	}

	// harusnya value dari counter adalah 2000 * 500 = 1.000.000
	wg.Wait()
	log.Println("counter : ", counter)
}
