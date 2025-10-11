package main

import "fmt"

func main() {
    // Variabel
    var nama string = "Achmad"
    umur := 25
    tinggi := 170.5
    aktif := true

    // Konstanta
    const negara = "Indonesia"
	var unused = "Halo"
    _ = unused  // tanda bahwa kita sengaja tidak pakai variabel ini

    // Output
    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Tinggi:", tinggi)
    fmt.Println("Aktif:", aktif)
    fmt.Println("Negara:", negara)
}
