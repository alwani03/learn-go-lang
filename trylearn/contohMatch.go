package main

import "fmt"

func main() {
	var angka1 int = 8
	var angka2 int = 3
	var hasil float64

	fmt.Println("Penjumlahan:", angka1+angka2)
	fmt.Println("Pengurangan:", angka1-angka2)
	fmt.Println("Perkalian:", angka1*angka2)
	fmt.Println("Sisa Bagi:", angka1%angka2)

	// Pembagian float
	hasil = float64(angka1) / float64(angka2)
	fmt.Println("Pembagian (float):", hasil)
}
