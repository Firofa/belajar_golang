package main

import "fmt"

func main() {
	var job string

	job = "Web Developer"

	fmt.Println("My Job is ", job)

	job = "Programmer"

	fmt.Println("My Job is ", job)

	// Sebenarnya tidak perlu menjelaskan tipe data apa ke si golang karena dia udah pinter, asalkan variable nya langsung di isi value :v

	var favouriteNumber = 23 //Jika ditulis begini tipe data akan otomatis 'int' dengan minimum 32 bit atau 64 bit sesuai dengan OS

	fmt.Println("My favourite number is ", favouriteNumber)

	var secondFavouriteNumber int8 = 19 //Gunakan Seperti ini jika ingin menggunakan int8 atau jenis tipe data number yg lebih spesifik lainnya

	fmt.Println("My second favourite number is ", secondFavouriteNumber)

	// Sebenarnya menggunakan 'var' ketika membuat variable tidak wajib asalkan menggunakan := ketika awal inisiasi variable

	nickname := "Firofa"

	fmt.Println("My Nickname is ", nickname)

	// Ketika menggunakan := untuk pemanggilan berikutnya ganti menggunakan =

	nickname = "NotFirofa"

	fmt.Println("My Nickname is Not ", nickname)
}
