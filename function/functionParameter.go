package main

import "fmt"

// Data dari luar yang dibutuhkan function disebut parameter
// Parameter itu opsional alias tidak wajib
// Namun jika menambahkan parameter di function, WAJIB MENGISI DATA PADA PARAMETER TERSEBUT SAAT MEMANGGIL FUNCTION NYA

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Ali", "Akmalughina")
}
