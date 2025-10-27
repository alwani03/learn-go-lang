package main

import "fmt"

func main() {
	var data = [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(data); i++ {
		fmt.Println("Angka ke -", i, ":", data[i])
	}

}
