package main

import "fmt"

// Type Declaration berguna untuk membuat tipe data alias dari tipe data sebenarnya

func main() {
	type NIP string   //NIP sebagai alias dari string
	type Married bool //Married sebagai alias dari boolean

	var Pegawai NIP = "123452313535346"
	const isMarried Married = false

	fmt.Println(Pegawai)
	fmt.Println(isMarried)
}
