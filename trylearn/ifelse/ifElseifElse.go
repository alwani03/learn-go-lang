package main

import "fmt"

func main() {
	var nilai int = 85

	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 75 {
		fmt.Println("Nilai B")
	} else {
		fmt.Println("Nilai C")
	}
}
