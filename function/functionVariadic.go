package main

import "fmt"

// Variadic Function
// Function yang memiliki Variable Argument
// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap aja semacam slice

// Apa bedanya dengan parameter biasa dengan tipe data array?
// -> Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// -> Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma.

func sumAll(numbers ...int) int { // parameter Variable Argument hanya bisa 1 dan itu juga harus paling kanan / final argument
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 20, 30, 40, 50)
	fmt.Println(total)

	slice := []int{12, 34, 24, 23, 33, 44}
	total2 := sumAll(slice...)
	fmt.Println(total2)

}
