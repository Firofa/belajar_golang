package main

import "fmt"

/**

Type Assertion
- Type Assertion merupakan kemampuan merubah tipe data menjadi tipe data yang di inginkan
- Fitur ini sering sekali digunakan ketika bertemu dengan data interface kosong

*/

func Random() interface{} {
	return "OK"
}

func main() {
	// Worked!!
	// result := Random()              //Tipe data masih interface kosong
	// resultString := result.(string) //Tipe data di ubah menjadi String
	// fmt.Println(resultString)

	// Not Worked!!
	// resultInt := result.(int) // panic (Bakal Error karena return dari interface kosong adalah string bukan int)
	// fmt.Println(resultInt)

	// Type Assertion Menggunakan Switch
	// Saat menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	// Jika panic dan tidak ter recover, maka otomatis program kita akan mati
	// Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

	// Worked and More Safe
	result := Random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

}
