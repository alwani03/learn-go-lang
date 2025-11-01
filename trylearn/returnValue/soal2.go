package main

import "fmt"

func totalBelanja(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}

func main() {
	fmt.Println(totalBelanja(10000, 20000, 5000))
}
