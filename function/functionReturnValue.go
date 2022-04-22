package main

import "fmt"

func getHello(firstName string) string { // Ketika membuat function dengan return jangan lupakan tipe data function nya
	return "Hello" + firstName
}

func main() {
	result := getHello("Eka")
	fmt.Println(result)
}
