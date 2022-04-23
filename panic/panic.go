package main

import "fmt"

// Panic function adalah function yang bisa kita gunakan untuk MENGHENTIKAN PROGRAM
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun DEFER FUNCTION TETAP AKAN DI EKSEKUSI

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
}

func main() {
	runApp(true)
}
