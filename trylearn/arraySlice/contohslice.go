// package main

// import "fmt"

// func main() {
// 	angka := []int{10, 20, 30} // slice
// 	fmt.Println(angka)

// 	// Menambahkan elemen
// 	angka = append(angka, 40, 50)
// 	fmt.Println(angka) // [10 20 30 40 50]
// }

package main

import "fmt"

func main() {
	data := []int{10, 20, 30, 40, 50}
	potongan := data[1:3] // ambil elemen index 1 dan 2
	fmt.Println(potongan) // [20 30]

	potongan[0] = 999
	fmt.Println(data) // [10 999 30 40 50] -> ikut berubah!
}
