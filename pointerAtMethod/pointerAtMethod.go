package main

import "fmt"

/**

---------------------
| Pointer di Method |
---------------------

- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu di duplikasi ketika memanggil method

*/

type Man struct {
	Name string
}

func (man *Man) Married() { // Menggunakan Pointer menghemat memory karena tidak menduplikat data
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di Method ", man.Name)
}

func main() {
	adi := Man{"Adi"}
	adi.Married()

	fmt.Println(adi.Name)

}
