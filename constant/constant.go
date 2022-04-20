package main

import "fmt"

// Constant itu mirip Variable cuman sekali di Inisialisasi gk bakal bisa di ubah lagi

// Constant Nilai nya gk perlu ditampilkan langsung di function itu gk kayak variable \\

func main() {

	const TypeOfPet string = "Cat"

	fmt.Println("My Type of Pet is ", TypeOfPet)

	const nameOfPet string = "Garfield" // Walaupun nameOfPet tidak di println sama sekali tapi golang tidak memberikan notifikasi error saat menggunakan constant berbeda dengan saat menggunakan variable

	// Multiple Const
	const (
		ageOfMyPet           int8 = 1
		favouriteFoodOfMyPet      = "Fish"
	)

	fmt.Println("The age of my Pet is ", ageOfMyPet)
	fmt.Println("My Pet's Favourite Food is ", favouriteFoodOfMyPet)

}
