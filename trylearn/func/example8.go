package main

import (
	"fmt"
	"strings"
)

func CountVowels(s string) int {
	s = strings.ToLower(s)
	count := 0
	for _, r := range s {
		switch r {
		case 'a', 'i', 'u', 'e', 'o':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountVowels("Belajar Go")) // output: 4
}
