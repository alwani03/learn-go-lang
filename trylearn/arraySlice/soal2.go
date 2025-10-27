package main

import "fmt"

func main() {
	angka := [5]int{5, 10, 15, 20, 25}

	total := 0
	for i := 0; i < len(angka); i++ {
		total += angka[i]
	}

	fmt.Println("Total:", total) // Total: 75

}
