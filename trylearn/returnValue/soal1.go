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
	pasangan := [][2]float64{{4, 2}, {10, 0}}
	var data []float64

	for _, p := range pasangan {
		hasil, err := bagi(p[0], p[1])
		if err != nil {
			fmt.Println("Error:", err)
			continue
		} else {
			fmt.Println("Hasil:", hasil)
		}
		data = append(data, hasil)
	}

}
