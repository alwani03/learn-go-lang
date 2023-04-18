package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpAlwani NoKtp = "312314121221212847"
	var status Married = true

	fmt.Println(noKtpAlwani)
	fmt.Println(status)

}
