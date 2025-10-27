// example 1,2,3,4
package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}

func split(n int) (a, b int) {
	a = n / 2
	b = n - a
	return a, b // mengembalikan a, b
}

func bagi(dividen, pembagi float64) (float64, error) {
	if pembagi == 0 {
		return 0, fmt.Errorf("pembagi nol")
	}
	return dividen / pembagi, nil
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println("Hasil tambah:", tambah(3, 4)) // 7
	result, err := bagi(3, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil bagi:", result) // 5

		a, b := split(8)
		fmt.Println("Hasil split:", a, b) // 4 5
	}

	fmt.Println("Hasil sum:", sum(1, 2, 3, 4, 5)) // 15
}
