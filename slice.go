package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "alwani")
	fmt.Println("----1", slice3)
	slice3[1] = "BukanDesember"
	fmt.Println("----2", slice3)

	/*
		*hati" dalam merubah array maka akan merubah slice nya*
		month[5] = "Diubah"
		fmt.Println(slice1)

		*hati" dalam merubah slice maka akan merubah slice nya*
		slice1[0] = "MeiDiubah"
		fmt.Println("--->", month)
	*/

}
