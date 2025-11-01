package main

import "fmt"

// Fungsi untuk menghitung rata-rata
func hitungRataRata(nilai []float64) float64 {
	total := 0.0
	for _, n := range nilai {
		total += n
	}
	return total / float64(len(nilai))
}

func main() {
	// Data nilai siswa
	nilaiSiswa := []float64{80, 75, 90, 85, 60}

	// Panggil fungsi
	rata := hitungRataRata(nilaiSiswa)

	fmt.Println("Data nilai:", nilaiSiswa)
	fmt.Printf("Rata-rata: %.2f\n", rata)
}
