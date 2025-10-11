package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b float64
	var op string

	fmt.Println("Kalkulator Mini")
	fmt.Println("Masukkan format: <angka1> <operator> <angka2>")
	fmt.Println("Contoh: 10 + 3")
	fmt.Print("> ")

	// Baca input; contoh: 10 + 3
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}

	switch op {
	case "+":
		fmt.Printf("%.6g %s %.6g = %.6g\n", a, op, b, a+b)
	case "-":
		fmt.Printf("%.6g %s %.6g = %.6g\n", a, op, b, a-b)
	case "*":
		fmt.Printf("%.6g %s %.6g = %.6g\n", a, op, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
			return
		}
		fmt.Printf("%.6g %s %.6g = %.6g\n", a, op, b, a/b)
	case "%":
		// modulo hanya untuk integer -> casting
		ai := int64(a)
		bi := int64(b)
		if bi == 0 {
			fmt.Println("Error: modulo dengan nol")
			return
		}
		fmt.Printf("%d %s %d = %d\n", ai, op, bi, ai%bi)
	default:
		fmt.Println("Operator tidak dikenal. Gunakan + - * / %")
	}
}
