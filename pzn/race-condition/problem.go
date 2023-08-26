package main

import "log"

func Problem() {
	// race condition terjadi ketika ada beberapa
	// goroutine yang mengubah data pada waktu yang sama
	// sehingga hanya dianggap satu operasi

	// saya akan membuat variable yang bisa
	// dipakai oleh semua goroutine
	data := 0
	// saya akan buat 1000 go routine
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go func() {
			defer wg.Done()
			// dimana setiap goroutine
			// akan mengubah variable data
			// menjadi 100
			for j := 1; j <= 100; j++ {
				data++
			}
		}()
	}
	wg.Wait()

	// harusnya variable data berjumlah 1000 * 100 = 100000
	// ternyata varible data hanya berjumlah sekitar 88000
	// artinya ada 12000 goroutine yang menaikan variable data
	// secara bersamaan, dan inilah yang dinamakan dengan race condition
	log.Println("data : ", data)
}
