package main

import "fmt"

// Recursive function adalah function yang memanggil function dirinya sendiri
// Kadang dalam pekerjaan, 	kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// Contoh kasus yang lebih mudah diselesaikan menggunakan recursive function adalah factorial

// Solusi looping:
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
