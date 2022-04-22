package main

import "fmt"

// Function tidak hanya bisa kita simpan didalam variable sebagai value
// Namun juga bisa kita gunakan sebagai parameter untuk function lain

// Type Declaration \/
type Filter func(string) string // dengan menggunakan type declaration pada function as parameter membuat kita tidak perlu repot ketika function membutuhkan parameter tambahan tinggal tambah saja e.g func(string, string, string)
// Type Declaration /\

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Kucing", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
