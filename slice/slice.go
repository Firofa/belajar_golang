package main

import "fmt"

// Slice adalah potongan dari data Array
// Slice mirip dengan array yang membedakan adalah ukuran slice bisa berubah
// Slice dan Array selalu terkoneksi, Dimana slice adalah data yang mengakses sebagian atau seluruh data array

// Pointer adalah penunjuk data pertama di array para slice
// Length adalah panjang dari slice
// Capacity adalah kapasitas dari slice, length tidak boleh lebih dari capacity

// Membuat Slice dari Array ada beberapa cara
// array[low:high] => Membuat slice dari array dimulai index low sampai dengan index sebelum high
// array[low:] 	=> Membuat slice dari array dimulai index low sampai dengan index akhir di array
// array[:high] => Membuat slice dari array dimulai index 0 sampai index sebelum high
// array[:]     => Membuat slice dari array dimulai index - sampai index akhir di array

// Function Slice
// len(slice) => Untuk Mendapatkan Length Slice
// cap(slice) => Untuk Mendapatkan Capacity Slice
// append(slice, data) => Membuat Slice baru dengan menambah data ke posisi terakhir slice, Jika capacity sudah penuh, maka akan membuat array baru
// make([]TypeData, length, capacity) => Membuat Slice Baru
// copy(destination, source) => Menyalin slice dari source ke destination

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //Untuk Melihat Length
	fmt.Println(cap(slice1)) //Untuk Melihat Capacity

	// months[5] = "Diubah"
	// fmt.Println(slice1) //Ketika array diubah slice akan ikut berubah

	// slice1[0] = "Bukan Mei"
	// fmt.Println(months) //Ketika slice diubah array akan ikut berubah

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) //Output: [Senin, Selasa, Rabu, Kamis, Jum'at, Sabtu Baru, Minggu Baru ]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) //Output: [Ups, Minggu Baru, Libur Baru] | Info: daySlice2 merupakan hasil append dari daySlice1 ditambah satu nilai "Libur Baru". Namun capacity daySlice1 sudah digunakan semua sehingga dibuat array baru untuk daySlice2
	fmt.Println(days)      //Output: [Senin, Selasa, Rabu, Kamis, Jum'at, Sabtu Baru, Minggu Baru] | Info: array days tidak berubah sewaktu ada append baru di daySlice2 karena slice sudah melebihi capacity sehingga daySlice2 membuat array baru

	// COba Function Make Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "web"
	newSlice[1] = "developer"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Coba Function Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) //Ketika melakukan copy pastikan length dan capacity dari slice source dan slice destionation nya sama agar data tidak terpotong

	fmt.Println(toSlice)

	// ! Hati - hati membedakan Array dan Slice
	iniArray := [...]int{1, 2, 3, 4, 5} // * Ketika Membuat array tambahkan "..." atau angka pada kurung []
	iniSlice := []int{1, 2, 3, 4, 5}    //* Ketika Membuat slice hanya memakai kurung []

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
