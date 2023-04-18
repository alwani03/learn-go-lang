package main

import "fmt"

func main() {
	var name string
	name = "achmad Alwani"
	fmt.Println(name)

	name = "Nurul Karomah"
	fmt.Println(name)

	fmt.Println("---------")

	var nicknamea = "Alwani"
	fmt.Println(nicknamea)

	var nicknameb = "Nurul"
	fmt.Println(nicknameb)

	fmt.Println("----------------------------------------------")
	var age = 23
	fmt.Println("Umur ", nicknamea, " dan ", nicknameb, " Adalah ", age, " Tahun")
	fmt.Println("----------------------------------------------")

	pendidikan := "S1"

	fmt.Println("Pendidikan ", nicknamea, " dan ", nicknameb, " Adalah ", pendidikan)
	fmt.Println("----------------------------------------------")

	var (
		jenis_kelamin1 = "laki-laki"
		jenis_kelamin2 = "Perempuan"
	)
	fmt.Println("Jenis Kelamin ", nicknamea, " Adalah ", jenis_kelamin1)
	fmt.Println("Jenis Kelamin ", nicknameb, " Adalah ", jenis_kelamin2)

}
