package main

import (
	"fmt"
	"log"
)

func main() {
	var umur int
	fmt.Print("Masukkan umur: ")
	_, err := fmt.Scanln(&umur)

	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	if umur < 13 {
		fmt.Println("Kamu anak-anak")
	} else if umur >= 13 && umur <= 17 {
		fmt.Println("Kamu remaja")
	} else if umur >= 18 && umur < 60 {
		fmt.Println("Kamu dewasa")
	} else {
		fmt.Println("Kamu lansia")
	}

}
