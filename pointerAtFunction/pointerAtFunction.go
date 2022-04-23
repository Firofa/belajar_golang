package main

import "fmt"

/**

------------------------
| Pointer di Function  |
------------------------

- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika kita ingin mengubah data didalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

*/

type Address struct {
	City, Province, Country string
}

// func ChangeAddressToIndonesia(address Address) { // (1)
// 	address.Country = "Indonesia"
// }

// func ChangeAddressToIndonesia(address *Address) { // (2) Berubah karena parameter nya menjadi pointer menggunakan operator &
// 	address.Country = "Indonesia"
// }

func ChangeAddressToIndonesia(address *Address) { // Cara (3)
	address.Country = "Indonesia"
}

func main() {
	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  ""}

	// ChangeAddressToIndonesia(alamat) // (1) Tidak berubah karena data duplikatnya yang dikirim
	// ChangeAddressToIndonesia(&alamat) // (2) Berubah karena telah menjadi pointer menggunakan operator &

	// Cara (3)
	var alamatPointer *Address = &alamat

	ChangeAddressToIndonesia(alamatPointer)

	fmt.Println(alamatPointer)

}
