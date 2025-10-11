package main

import "fmt"

func main() {

	// TODO: deklarasikan variabel di sini
	var nama string
	var umur int
	var tinggiBadan float64
	var statusAktif bool

	nama = "Alwani"
	umur = 26
	tinggiBadan = 157.6
	statusAktif = true

	const Negara = "Indonesia"

	fmt.Println("nama :", nama)
	fmt.Println("umur :", umur)
	fmt.Println("tinggiBadan :", tinggiBadan)
	fmt.Println("statusAktif :", statusAktif)
	fmt.Println("Negara :", Negara)

	fmt.Printf("Halo, nama saya %s. Saya berumur %d tahun dan tinggal di %s.\n", nama, umur, Negara)
	fmt.Printf("Status aktif saya: %t\n", statusAktif)

}
