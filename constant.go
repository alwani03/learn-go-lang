package main

import (
	"fmt"
)

func main() {
	const firstName string = "Achmad"
	const lastName = "Alwani"

	/* error */

	/*
		firstName= "test"
		lastName = "testubah"
	*/

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println("-----------")

	const (
		firstName2 = "Nurul"
		lastName2  = "Karomah"
	)

	fmt.Println(firstName2)
	fmt.Println(lastName2)

}
