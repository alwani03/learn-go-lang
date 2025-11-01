package main

import "fmt"

func hitung(angka ...int) (int, float64) {
	total := 0
	for _, v := range angka {
		total += v
	}
	rata := float64(total) / float64(len(angka))
	fmt.Println("Debug: Total di dalam fungsi hitung:", len(angka))
	return total, rata
}

func main() {
	total, rata := hitung(10, 20, 30, 40, 70)
	fmt.Println("Total:", total)
	fmt.Println("Rata-rata:", rata)
}
