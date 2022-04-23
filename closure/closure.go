package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// Gunakan fitur closure dengan bijak saat membuat aplikasi
// Karena jika salah menggunakan closure akan mengakibatkan terubah data yang tidak diinginkan

func main() {
	job := "Gamers"
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		// counter := 0
		counter++ //Perintah ini mengubah value dari counter di scope diatasnya (Closure) karena counter tidak di inisiasi di dalam scope sehingga mengambil inisiasi diatasnya
		// fmt.Println(counter)

		job := "Programmer" //Terkecuali jika ada inisiasi variable lagi di dalam scope sehingga variable diluar scope tidak akan terpengaruh dengan perubahan pada variable didalam scope
		fmt.Println(job)
	}

	increment()

	fmt.Println(counter)
	fmt.Println(job)
}
