package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "Belajar Go"
	replace := strings.ReplaceAll(kalimat, " ", "")

	total_vocal := 0

	for index, char := range replace {
		if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' || char == 'A' || char == 'I' || char == 'U' || char == 'E' || char == 'O' {
			total_vocal++
		}
		fmt.Printf("Index ke-%d : %c\n", index, char)
	}

	fmt.Printf("\nTotal huruf vocal: %d\n", total_vocal)
}
