package main

import (
	"errors"
	"fmt"
)

/**

Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

Membuat Error Interface
- Untuk Membuat Error, Kita tidak perlu manual
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors

*/

// type error interface {
// 	Error() error,
// }

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// var contohError error = errors.New("ups, Error")
	// fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err)
	}
}
