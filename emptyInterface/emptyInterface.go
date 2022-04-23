package main

import "fmt"

/**
----------------------------------
Interface Kosong / Empty Interface
----------------------------------
- Go-lang bukanlah bahasa pemrograman yang berorientasi objek
- Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- Contoh di Java ada java.lang.Object
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
- Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

Contoh Interface Kosong:
type Apapun interface {} //Kontrak yang didalamnya tidak ada isi kontraknya

*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups.."
	}
}

func sayUwuu(i interface{}) interface{} {
	return "Uwuu"
}

func main() {
	var dataOne interface{} = Ups(1)
	var dataTwo interface{} = Ups(2)
	var dataThree interface{} = Ups(3)

	fmt.Println(dataOne)
	fmt.Println(dataTwo)
	fmt.Println(dataThree)

	var personOne interface{} = sayUwuu(nil)

	fmt.Println(personOne)
}
