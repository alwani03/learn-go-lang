package main

import "fmt"

func main() {

	k := manusia{tubuh: "kaki"}
	t := manusia{tubuh: "tangan"}

	k.berjalan()
	t.makan()

}

type manusia struct {
	tubuh string
}

func (m manusia) berjalan() {
	fmt.Println("bergerak bagian tubuh ", m.tubuh)
}

func (m manusia) makan() {
	fmt.Println("bergerak bagian tubuh ", m.tubuh)
}
