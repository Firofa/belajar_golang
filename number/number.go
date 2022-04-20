package main

import "fmt"

// Untuk Build Gunakan: go build number.go << Akan Menghasilkan number.exe >>>
// Untuk Mode Development Gunakan: go run number.go

// Tipe Data Number
// int8 (Integer 8) -> Nilai Minimum: -128 (Minus seratusan) & Nilai Maksimum: 127 (Seratusan)
// int16 -> Nilai Minimum: -32768 (Minus tiga puluh dua ribuan) & Nilai Maksimum: 32767 (Tiga puluh dua ribuan)
// int32 -> Nilai Minimum: -2147483648 (Minus dua milyaran) & Nilai Maksimum: 2147483647 (Dua Milyaran)
// int64 -> Nilai Minimum: -9223372036854775808 (Minus Banyak :v) & NIlai Maksimum: 9223372036854775807 (Banyak Pokoknya :v)

// uint8 (Unsigned Integer 8) -> Nilai Minimum 0 & Nilai Maksimum 255 ( Minimal 0 sampe 255, Maksimal dua ratus lima puluhan, Gk ada nilai negatif, Minus di int8 ditambahin ke positif :3 )
// uint16 -> Nilai Minimum 0 & Nilai Maksimum 65535 (Minimal 0 sampe maksimal enam puluh lima ribuan..)
// uint32 -> Nilai Minimum 0 & Nilai Maksimum 4294967295 (Minimal 0 sampe maksimal empat milyaran..)
// uint64 -> Nilai Minimum 0 & Nilai Maksimum 18446744073709551615 (Minimal 0 sampe maksimal banyak pokoknya :v )

// Tipe Data Floating Point (Sejenis bentuk-bentuk desimal/float)
// float32 -> Nilai Minimum 1.18 x 10^-38 & Nilai Maksimum 3.4 x 10^38 (Sering Digunakan)
// float64 -> Nilai Minimum 2.23 x 10^-308 & Nilai Maksimum 1.80 x 10^308 (Sering Digunakan)
// complex64 -> complex numbers with float32 real and imaginary parts. (Jarang Digunakan kecuali buat yang Statistika, Hitung" an ala ala Matematikawan)
// complex128 -> complex numbers with float64 real and imaginary parts. (Jarang Digunakan kecuali buat yang Statistika, Hitung" an ala ala Matematikawan)

// Tipe Data Alias (Nama Samaran si tipe data (maybe...))
// byte -> uint8
// rune -> int32
// int  -> Minimal int32 atau Minimal int64 (Menyesuaikan dengan sistem operasi 32 bit atau 64 bit)
// uint -> Minimal int32 atau Minimal int64 (Menyesuaikan dengan  sistem operasi 32 bit atau 64 bit)

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}
