package main

import "fmt"

func statistik(angka ...int) (int, float64) {
	total := 0

	if len(angka) == 0 {
		return 0, 0.0

	}

	for _, v := range angka {
		total += v
	}

	rataRata := float64(total) / float64(len(angka))

	return total, rataRata
}

func main() {

	total, rata := statistik(50, 50, 25, 25)
	total2, rata2 := statistik(0)

	fmt.Println("Total:", total)
	fmt.Println("Rata-rata:", rata, "\n--------------------")

	fmt.Println("Total:", total2)
	fmt.Println("Rata-rata:", rata2)

}
