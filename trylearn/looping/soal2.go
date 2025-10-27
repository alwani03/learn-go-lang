package main

import "fmt"

func main() {

	barang := map[string]int{
		"Kaos":   5,
		"Celana": 3,
		"Topi":   2,
	}

	total := 0

	for val, jumlah := range barang {

		fmt.Printf("Nama Barang: %s\n", val)
		total += jumlah
		if jumlah > 2 {
			fmt.Printf("Jumlah barang %s lebih dari 2\n", val)
		}
	}

	fmt.Printf("Total seluruh barang: %d\n", total)

}
