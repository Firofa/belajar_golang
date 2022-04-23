package main

import "fmt"

/**
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
- Sebuah Interface berisikan definisi - definisi method
- Biasanya Interface digunakan sebagai kontrak
*/

/**
Implementasi Interface

- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- Sehingga kita tidak perlu mengimplementasikan interface secara manual
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
*/

type HasName interface { // (1) Tipe data apapun yang memiliki function getName dan return value nya string maka berhak mengikuti kontrak ini
	GetName() string
}

func SayHello(hasName HasName) { // (4) Tipe data yang kontraknya sesuai dengan interface HasName dapat menggunakan function SayHello()
	fmt.Println("Hello", hasName.GetName())
}

type Person struct { // (2) Tipe data struct ini...
	Name string
}

func (person Person) GetName() string { // (3) ...Memiliki function GetName dengan return string sehingga berhak mengikuti kontrak interface HasName
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person   // (5) Menginisiasi sebuah variable ber tipe data struct
	eko.Name = "Eko" // (6) Memberikan value "Eko" pada variable eko.Name

	SayHello(eko) // (7) Memanggil Function SayHello dengan parameter yang sesuai dengan kontrak HasName Interface sehingga tidak error

	cat := Animal{
		Name: "Miaw",
	}

	SayHello(cat)
}
