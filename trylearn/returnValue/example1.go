package main

import "fmt"

// Fungsi mengembalikan dua nilai: string dan int
func getProfile() (string, int) {
	name := "Achmad"
	age := 26
	return name, age
}

func main() {
	nama, umur := getProfile()
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
}
