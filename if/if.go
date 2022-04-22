package main

import "fmt"

func main() {
	name := "Web Dev"

	if name == "Web Dev" {
		fmt.Println("Hello", name)
	} else if name == "Android Dev" {
		fmt.Println("Hai", name)
	} else {
		fmt.Println("Who are you?")
	}

	// If dengan short statement
	// * If mendukung short statement sebelum kondisi
	// * Hal ini cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu Panjang")
	}
}
