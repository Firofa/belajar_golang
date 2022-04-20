package main

import "fmt"

// string
// Kata Kunci: string (Semua Huruf Kecil)
// Diawali dengan " dan diakhiri dengan "

func main() {
	fmt.Println("Fi")
	fmt.Println("Fi Ro")
	fmt.Println("Fi Ro Fa")

	// Function untuk menghitung jumlah karakter menggunakan len()
	fmt.Println(len("Hello World"))

	// Function untuk mendapatkan nilai byte dari sebuah karakter pada string menggunakan "string"[number]
	fmt.Println("Hello World"[0]) //Mengambil byte karakter 'H'
	fmt.Println("Hello World"[6]) //Mengambil byte karakter 'W'
}
