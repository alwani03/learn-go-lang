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

	fmt.Println("\n=== HASIL PERBANDINGAN ===")
	fmt.Printf("%.2f == %.2f → %v\n", a, b, a == b)
	fmt.Printf("%.2f != %.2f → %v\n", a, b, a != b)
	fmt.Printf("%.2f >  %.2f → %v\n", a, b, a > b)
	fmt.Printf("%.2f <  %.2f → %v\n", a, b, a < b)
	fmt.Printf("%.2f >= %.2f → %v\n", a, b, a >= b)
	fmt.Printf("%.2f <= %.2f → %v\n", a, b, a <= b)
}
