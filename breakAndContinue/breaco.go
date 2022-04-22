package main

import "fmt"

// break digunakan untuk menghentikan seluruh perulangan
// continue digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya

func main() {

	// * Kode Program Break
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke", i)
	// }

	// * Kode Program Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}

}
