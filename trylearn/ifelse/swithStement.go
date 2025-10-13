package main

import "fmt"

func main() {
	warna := "merah"

	switch warna {
	case "merah":
		fmt.Println("Berhenti!")
	case "kuning":
		fmt.Println("Hati-hati!")
	case "hijau":
		fmt.Println("Jalan!")
	default:
		fmt.Println("Warna tidak dikenal")
	}
}
