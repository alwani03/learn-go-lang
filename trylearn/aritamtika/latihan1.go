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

	fmt.Println("-----------------------------")
	fmt.Println("Hasil Penjumlahan =", a+b)
}
