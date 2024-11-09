package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "alwani"
	names[1] = "nurul"
	names[2] = "kharomah"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	//command

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println("---------")
	fmt.Println(values)
	fmt.Println("---------")

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("---------")
	fmt.Println("---------")

	fmt.Println(len(names)) /*jumlah length*/

	var lagi [10]int
	fmt.Println(len(lagi)) /*jumlah length*/

}
