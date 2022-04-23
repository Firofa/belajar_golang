package main

import "fmt"

// Secara Default di Go-Lang semua variable itu di passing by value, bukan by reference
// Artinya Jika kita mengirim sebuah variable ke dalam function, method atau variable lain sebenarnya yang dikirim adalah duplikasi value nya

/**
-----------
- POINTER -
-----------
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

 -------------
 | OPERATOR & |
 -------------
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya

--------------
| OPERATOR * |
--------------
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

---------------
|Function New |
---------------
- Sebelumnya untuk membuat pointer dengan menggunakan Operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

*/

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 Address = address1 // Secara Default address2 akan menerima pass by value dari address1

	var address3 *Address = &address1 // Menggunakan operator & untuk pass by reference, sehingga bisa dibilang address3 itu pointer ke address1 / *address3

	address2.City = "Bandung"

	address3.City = "Surabaya"

	address3 = &Address{"Malang", "Jawa Timur", "Indonesia"} //Mengubah pointer address3 yang sebelumnya pointer ke address1, berubah menjadi pointer ke &Address{"Malang", "Jawa Timur", "Indonesia"} tanpa mengajak address1

	*address3 = Address{"Magelang", "Jawa Tengah", "Indonesia"} //Mengubah pointer address3 yang sebelumnya pointer ke address1, berubah menjadi pointer ke Address{"Magelang", "Jawa Tengah", "Indonesia"} serta mengajak address1 dan siapun yang mengacu pada address1 menjadi pointer ke Address{"Magelang", "Jawa Tengah", "Indonesia"}

	var address4 *Address = new(Address) // dengan function new, address4 akan menjadi pointer ke data kosong
	address4.City = "Pontianak"
	fmt.Println(address4) // Output: &{Pontianak }

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
