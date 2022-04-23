package main

import "fmt"

/**
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan didalam field
- Sederhananya struct adalah kumpulan dari field
*/

/**
Membuat Data Struct

- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat

*/

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var person1 Customer
	person1.Name = "Alif"
	person1.Address = "Indonesia"
	person1.Age = 19

	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person1.Address)
	fmt.Println(person1.Age)

	// Penulisan Struct Literal
	person2 := Customer{
		Name:    "Didi",
		Address: "Malaysia",
		Age:     20,
	}

	person3 := Customer{"Mail", "Malaysia", 12} // Cara ini lebih singkat tapi ketika ada field baru di struct akan ERROR (Tidak disarankan Menggunakan cara ini)

	fmt.Println(person2)
	fmt.Println(person3)
}
