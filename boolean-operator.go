package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsesnsi = absensi >= 80

	var lulus = lulusAbsesnsi && lulusUjian

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsesnsi)

	fmt.Println("----------|")
	fmt.Println(lulus, "	  |")
	fmt.Println("----------|")

	/* atau kita bisa menggunakan cara sebegai berikut */

	fmt.Println(ujian >= 80 && absensi >= 80)

}
