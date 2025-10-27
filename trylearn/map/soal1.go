package main

import "fmt"

func main() {
	// 1. Buat map laci

	laci := map[string]int{
		"Kaos":   5,
		"Celana": 3,
		"Topi":   2,
	}

	// 2. Tambah data baru
	laci["Sepatu"] = 4
	// 3. Hapus data "Topi"
	delete(laci, "Topi")
	// 4. Loop tampilkan semua isi map

	var check_jaket bool = false

	for item, jumlah := range laci {
		fmt.Printf("Laci %s berisikan %d\n", item, jumlah)

		if item == "Jaket" {
			check_jaket = true
		}
	}

	// 5. Bonus: cek apakah "Jaket" ada
	if check_jaket {
		fmt.Println("Jaket ada di laci")
	} else {
		fmt.Println("Laci Jaket tidak ditemukan")
	}
}
