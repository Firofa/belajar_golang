package main

import "fmt"

/**
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function

- Namun jika kita ingin menambahkan method kedalam structs, sehingga seakan-akan sebuah struct memiliki function

- Method adalah function

*/

type Customer struct {
	Name    string
	Address string
	Age     int
}

// Function Biasa
// func sayHi(customer Customer, name string) {
// 	fmt.Println("Hello", name, "my Name is", customer.Name)
// }

// Struct Function
func (customer Customer) sayHi(name string) { //Ketika membuat function khusus untuk sebuah struct disarankan menggunakan struct function
	fmt.Println("Hello", name, "my Name is", customer.Name)
}

func (a Customer) sayYoho() {
	fmt.Println("Yohohohohoho~ From", a.Name)
}

func main() {
	personOne := Customer{
		Name:    "Alfin",
		Address: "Indonesia",
		Age:     22,
	}

	// Pemanggilan Function Biasa
	// sayHi(personOne, "Wiro")

	// Pemanggilan Struct Function
	personOne.sayHi("Wiro")
	personOne.sayYoho()
}
