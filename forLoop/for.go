package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// * For dengan statement
	// Init Statement -> Statement sebelum for di eksekusi
	// Post Statement -> Statement yang selalu di eksekusi setiap perulangan selesai
	for i := 0; i <= 10; i++ {
		fmt.Println("Nilai i ke:", i)
	}

	// * For Range
	// For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contohnya Array, Slice dan Map

	// slice := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	// for j := 0; j < len(slice); j++ {
	// 	fmt.Println(slice[j])
	// }

	// Kode diatas menggunakan kode for range menjadi
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	for index, day := range days { //Comment: Mirip foreach :v
		fmt.Println("index", index, "=", day)
	}

	// Untuk Menggunakan sebuah variable yang tidak akan ditampilkan tapi menghindari error si golang gunakan _
	// Misal:
	months := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	for _, month := range months { // variable _ hanya di inisiasi tapi tidak digunakan namun golang tidak menganggap error :3
		fmt.Println(month)
	}

	// For Range untuk Map
	person := make(map[string]string)
	person["platform"] = "Web"
	person["job"] = "Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
