package main

import "fmt"

func main() {
	angka := []int{10, 20, 30, 40, 50}
	potongan := angka[1:3] // ambil elemen index 1 dan 2

	potongan[0] = 999

	fmt.Println("Potongan : ", potongan)
	fmt.Println("Data Asli : ", angka)

}
