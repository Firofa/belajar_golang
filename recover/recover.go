package main

import "fmt"

// Recover adalah function yang bisa digunakan untuk menangkap data panic
// Dengan Recover proses panic akan terhenti, Sehingga Program akan tetap berjalan

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover() //Menangkap data panic
	if message != nil {
		fmt.Println("Terjadi Error dengan message: ", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

}

func main() {
	runApp(false)
	fmt.Println("Done...")
}
