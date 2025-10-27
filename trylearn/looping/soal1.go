package main

import "fmt"

func main() {

	angka := []int{2, 4, 6, 8, 10}
	total := 0

	total_up_5 := 0

	for _, val := range angka {

		fmt.Printf("Value sebagai berikut : %d \n", val)
		total += val

		if val > 5 {
			total_up_5++
		}
	}

	fmt.Printf("Total penjumlahan: %d\n", total)
	fmt.Printf("Total diatas: %d\n", total_up_5)

}
