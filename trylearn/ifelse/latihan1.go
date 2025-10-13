package main

import (
	"fmt"
	"log"
)

func main() {
	var param int
	fmt.Print("Masukkan nomor hari (1-7) : ")
	_, err := fmt.Scanln(&param)

	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
		return
	}

	switch param {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Nomor hari tidak valid")
	}

}
