package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Menu:")
	fmt.Println("1. Nasi Goreng")
	fmt.Println("2. Mie Goreng")
	fmt.Println("3. Soto Ayam")
	fmt.Println("4. Keluar")

	var param int
	fmt.Print("Pilihlah Menu 1-4:  ")
	_, err := fmt.Scanln(&param)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
		return
	}

	switch param {
	case 1:
		fmt.Println("Anda memilih Nasi Goreng")
	case 2:
		fmt.Println("Anda memilih Mie Goreng")
	case 3:
		fmt.Println("Anda memilih Soto Ayam")
	case 4:
		fmt.Println("Terima kasih, sampai jumpa!")
	default:
		fmt.Println("Nomor Menu Tidak Valid")
	}

}
