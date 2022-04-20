package main

import "fmt"

func main() {
	var SebuahString string = "Ini String"

	var SebuahByte = SebuahString[0]

	var KarakterByte = string(SebuahByte)

	fmt.Println(KarakterByte)

	// Dalam Melakukan Konversi tipe data hati-hati ketika tipe data yang akan diubah kapasitas nya tidak mencukupi (Tidak akan ada error hanya saja nilai data nya akan berubah)
	var rekorPorsiMakanRendang int32 = 129
	var rekorPorsiMakanRendang64 int64 = int64(rekorPorsiMakanRendang)
	var rekorPorsiMakanRendang8 int8 = int8(rekorPorsiMakanRendang64) // int8 tidak akan cukup untuk menampung kapasitas angka 129 karena nilai minimum dari int8 adalah -128 dan nilai maksimumnya adalah 127

	fmt.Println(rekorPorsiMakanRendang)
	fmt.Println(rekorPorsiMakanRendang64)
	fmt.Println(rekorPorsiMakanRendang8) // Karena tidak menampung nilai 129 maka rekorPorsiMakanRendang8 akan menghitung hingga 127 kemudian kembali ke -128 -> -127 sehingga nilai nya menjadi -127

}
