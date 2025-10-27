// example 5, 6
package main

import "fmt"

func increment(x *int) {
	*x = *x + 1
}

func apply(x int, f func(int) int) int {
	return f(x)
}

func main() {
	v := 7
	increment(&v)                                // v sekarang 6
	fmt.Println("Nilai v setelah increment:", v) // Output: Nilai v setelah increment: 6

	double := func(n int) int { return n * 2 }
	double2 := func(n int) int { return n * 7 * 19 }

	fmt.Println(apply(5, double))  // 10
	fmt.Println(apply(5, double2)) // 10

}
