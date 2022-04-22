package main

import "fmt"

func getFullName() (string, string) {
	return "Alif", "Bob"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Jika ingin menghiraukan salah satu value gunakan _ untuk menghiraukan
	firstName2, _ := getFullName()
	fmt.Println(firstName2)

}
