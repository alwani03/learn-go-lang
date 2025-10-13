package main

import "fmt"

func main() {
	nilai := 80

	switch {
	case nilai >= 90:
		fmt.Println("Sangat Baik")
	case nilai >= 75:
		fmt.Println("Baik")
	default:
		fmt.Println("Cukup")
	}
}
