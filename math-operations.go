package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	fmt.Println("-------------------")

	var a = 3
	var b = 10
	var c = a + b

	fmt.Println("Hasilnya adalah", c)
	fmt.Println("-------------------")

	var i = 10
	i += 10

	fmt.Println("Hasilnya i adalah", i)
	fmt.Println("-------------------")

	var j = 1
	j++
	j++

	var positif = +100
	var negatif = -100

	fmt.Println("Print", j, positif, negatif)

}
