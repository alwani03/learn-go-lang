package main

import "fmt"

func main() {
	// 1. Buat slice of map

	gudang := []map[string]int{
		{
			"Kaos":   5,
			"Celana": 3,
		},
		{
			"Sepatu": 4,
			"Topi":   2,
		},
		{
			"Jaket": 6,
			"Kaos":  2,
		},
	}

	totalKaos := 0

	// 2. Loop semua lemari dan tampilkan isinya
	for i, lemari := range gudang {
		fmt.Println("Isi Lemari", i+1)
		for item, jumlah := range lemari {
			fmt.Printf("- %s: %d\n", item, jumlah)
			if item == "Kaos" {
				totalKaos += jumlah
			}
		}
		fmt.Println()
	}

	// 3. Bonus: hitung total barang "Kaos"

	fmt.Printf("Total Kaos di semua lemari: %d\n", totalKaos)
}
