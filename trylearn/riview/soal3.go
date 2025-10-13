package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Pilih Aritmatika Menu:")
		fmt.Println("1. Tambah")
		fmt.Println("2. Kurang")
		fmt.Println("3. Kali")
		fmt.Println("4. Bagi")
		fmt.Println("5. Keluar")

		var param int
		fmt.Print("Pilihlah Menu 1-5: ")
		_, err := fmt.Scanln(&param)
		if err != nil {
			fmt.Println("Input menu tidak valid. Masukkan angka 1-5.")
			// buang sisa input dari buffer dengan membaca newline:
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if param == 5 {
			fmt.Println("Terima kasih, sampai jumpa!")
			break
		}

		// Validasi pilihan menu sebelum meminta angka
		if param < 1 || param > 4 {
			fmt.Println("Nomor menu tidak valid. Silakan pilih 1-5.")
			continue
		}

		var a, b float64
		fmt.Print("Masukkan angka pertama: ")
		_, err = fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Input angka tidak valid. Coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		fmt.Print("Masukkan angka kedua: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Input angka tidak valid. Coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		fmt.Printf("\n=== HASIL OPERASI ARITMATIKA ===\n\n")
		switch param {
		case 1:
			fmt.Printf("Hasil Penjumlahan: %.6g + %.6g = %.6g\n", a, b, a+b)
		case 2:
			fmt.Printf("Hasil Pengurangan: %.6g - %.6g = %.6g\n", a, b, a-b)
		case 3:
			fmt.Printf("Hasil Perkalian: %.6g * %.6g = %.6g\n", a, b, a*b)
		case 4:
			if b == 0 {
				fmt.Println("Error: pembagian dengan nol")
				// jangan exit; kembali ke menu
			} else {
				fmt.Printf("Hasil Pembagian: %.6g / %.6g = %.6g\n", a, b, a/b)
			}
		}
		fmt.Printf("\n=====================\n\n")
	}
}
