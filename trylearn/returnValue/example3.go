package main

import "fmt"

func jumlahAngka(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}

func main() {
	fmt.Println(jumlahAngka(1, 2, 3))       // 6
	fmt.Println(jumlahAngka(5, 10, 15, 20)) // 50
}
