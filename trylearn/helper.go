package main

import (
    "fmt"
    "os"
)


func sayHello(name string) {
    fmt.Println("Halo,", name)
	os.Exit(1) // langsung keluar (kode 1 artinya error)
    fmt.Println("Tidak akan dieksekusi")
}
