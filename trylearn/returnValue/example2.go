package main

import (
	"errors"
	"fmt"
)

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}
}
