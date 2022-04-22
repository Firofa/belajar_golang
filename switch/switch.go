package main

import "fmt"

func main() {
	name := "Hiccup"

	switch name {
	case "Alfin":
		fmt.Println("Hi", name)
	case "Hiccup":
		fmt.Println("yo", name)
		fmt.Println("yo", name)
	default:
		fmt.Println("How are you?", name)
	}

	// * Swith dengan Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Panjang...")
	case false:
		fmt.Println("Pendek...")
	}

	// * Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
