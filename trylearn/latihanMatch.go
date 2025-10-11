package main

import "fmt"

func main() {

	var angka1 int
	var angka2 int
	var Penjumlahan int
	var Pengurangan int
	var Perkalian int
	var Pembagian float64
	var SisaBagi int

	angka1 = 10
	angka2 = 3

	Penjumlahan = angka1 + angka2
	Pengurangan = angka1 - angka2
	Perkalian = angka1 * angka2
	Pembagian = float64(angka1) / float64(angka2)
	SisaBagi = angka1 % angka2

	Penjumlahan += 1
	fmt.Println("Penjumlahan     : ", Penjumlahan)
	fmt.Println("Pengurangan     : ", Pengurangan)
	fmt.Println("Perkalian 	: ", Perkalian)
	fmt.Println("Pembagian 	: ", Pembagian)
	fmt.Println("SisaBagi 	: ", SisaBagi)

}
