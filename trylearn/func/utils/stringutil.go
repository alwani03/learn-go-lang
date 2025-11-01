package stringutil

import "strings"

// CountVowels menghitung jumlah vokal dalam string.
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
