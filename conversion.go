package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)
	fmt.Println("-------")

	/*Convert to String*/
	var name = "Nurul"
	var r byte = name[2]
	var rString string = string(r)

	fmt.Println(rString)

}
