package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b float64

	fmt.Print("Masukkan angka pertama: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	fmt.Print("Masukkan angka kedua: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	fmt.Println("\n=== HASIL OPERASI ARITMATIKA ===")
	fmt.Println("Penjumlahan   :", a+b)
	fmt.Println("Pengurangan   :", a-b)
	fmt.Println("Perkalian     :", a*b)

	if b == 0 {
		fmt.Println("Pembagian     : Error (pembagian dengan nol)")
		fmt.Println("Sisa Bagi     : Error (modulo dengan nol)")
	} else {
		fmt.Println("Pembagian     :", a/b)
		fmt.Println("Sisa Bagi     :", int64(a)%int64(b))
	}
}
