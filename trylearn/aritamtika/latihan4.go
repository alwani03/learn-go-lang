package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	var a, b float64
	var operator string

	fmt.Print("Masukkan angka pertama: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	fmt.Print("Masukkan operator (+ - * / %): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}
	operator = strings.TrimSpace(operator)

	fmt.Print("Masukkan angka kedua: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	switch operator {
	case "+":
		fmt.Printf("Hasil Penjumlahan:  %g\n", a+b)
	case "-":
		fmt.Printf("Hasil Pengurangan:  %g\n", a-b)
	case "*":
		fmt.Printf("Hasil Perkalian:    %g\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
			return
		}
		fmt.Printf("Hasil Pembagian:    %g\n", a/b)
	case "%":
		if b == 0 {
			fmt.Println("Error: modulo dengan nol")
			return
		}
		// math.Mod untuk modulo pada float64
		fmt.Printf("Sisa Bagi (mod) :   %g\n", math.Mod(a, b))
	default:
		fmt.Println("Operator tidak dikenal. Gunakan + - * / %")
	}
}
