package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	salinan := make([]int, len(data))

	copy(salinan, data)
	salinan[0] = 999

	fmt.Println("Data asli :", data)
	fmt.Println("Salinan	  :", salinan)

}
