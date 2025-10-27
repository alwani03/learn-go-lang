package main

import "fmt"

func main() {
	var angka [3]int // array berisi 3 elemen bertipe int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30

	fmt.Println(angka)      // [10 20 30]
	fmt.Println(len(angka)) // panjang array = 3
}

// package main

// import "fmt"

// func ubahArray(a [3]int) {
// 	a[0] = 999
// 	fmt.Println("Di dalam fungsi ubahArray:", a) // Di dalam fungsi ubahArray: [999 2 3]
// }

// func main() {
// 	data := [3]int{1, 2, 3}
// 	ubahArray(data)
// 	fmt.Println(data) // tetap [1 2 3], tidak berubah
// }
