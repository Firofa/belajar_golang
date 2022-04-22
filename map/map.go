package main

import "fmt"

// Pada Array dan Slice untuk mengakses data kita menggunakan index Number dimulai dari 0
// Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
// Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci-nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
// Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci yang sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

// Penulisan Map :
// var "nama_variable" map["tipe_data_key"]"tipe_data_value" = map["tipe_data_key"]"tipe_data_value" {"key": "value"}

// Atau Bisa Juga:
// var "nama_variable" map["tipe_data_key"]"tipe_data_value" = make(map[TypeKey]TypeValue)

// Atau Singkatnya :
//  "nama_variable" := map["tipe_data_key"]"tipe_data_value" {"key" : "value"}

// * Function pada Map:
// len(map) => Untuk Mendapatkan Length di map
// map[key] => Mengambil data di map dengan key
// map[key] = value => Mengubah data di map dengan key
// make(map[TypeKey]TypeValue) => Membuat Map Baru
// delete(map, key) => Menghapus data di map dengan key

func main() {

	person := map[string]string{
		"name":    "alfin",
		"address": "bandung",
	}

	// Untuk Menambah Data Map
	person["job"] = "Web Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Untuk Menghapus Data Map
	delete(person, "job")
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "How To Train Your Dragon"
	book["author"] = "Cressida Cowell"
	book["abcd"] = "efgh"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "abcd")
	fmt.Println(book)
	fmt.Println(len(book))

}
