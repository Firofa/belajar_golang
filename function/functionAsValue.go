package main

import "fmt"

// Function di Golang bisa dibuat sebagai tipe data, variable, independen tanpa grup, dsb
// Function adalah first class citizen

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye //Membuat function sebagai value
	fmt.Println(goodbye("Cat"))
}
