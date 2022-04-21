package main

import "fmt"

// Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
// Daya tampung array tidak bisa bertambah setelah Array dibuat

func main() {
	var names [3]string

	names[0] = "Hello"
	names[1] = "Programmer"
	names[2] = "World"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membyat Array Langsung
	var values = [5]int{
		23,
		19,
		24,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Memanipulasi & Mengakses Data Array
	// len(array) -> Digunakan untuk mendapatkan panjang Array

	fmt.Println(len(names)) // Mengecek berapa panjang array (bukan jumlah data)
	fmt.Println(len(values))
}
