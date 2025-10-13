package main

import (
	"fmt"
	"log"
)

func main() {

	var param int
	fmt.Print("Masukkan nilai:  ")
	_, err := fmt.Scanln(&param)
	if err != nil {
		log.Fatalf("Input tidak valid: %v\n", err)
	}
	switch {
	case param < 60:
		fmt.Println("D")
	case param >= 60 && param < 75:
		fmt.Println("C")
	case param >= 75 && param < 90:
		fmt.Println("B")
	case param >= 90 && param <= 100:
		fmt.Println("A")
	default:
		fmt.Println("Nilai tidak valid")

	}
}
