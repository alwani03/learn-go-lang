package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Kalkulator Mini ===")
	fmt.Println("Ketik 'exit' untuk keluar")
	fmt.Println("Format: <angka1> <operator> <angka2>")
	fmt.Println("Contoh: 10 + 3")
	fmt.Println("------------------------")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		// Baca input 1 baris penuh
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Gagal membaca input: %v\n", err)
		}

		// Hilangkan spasi dan newline
		input = strings.TrimSpace(input)

		// Cek apakah user ingin keluar
		if strings.ToLower(input) == "exit" {
			fmt.Println("Program selesai. Terima kasih 🙌")
			break
		}

		// Pisahkan input menjadi bagian-bagian
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Format salah! Gunakan format: <angka1> <operator> <angka2>")
			continue
		}

		// Ambil bagian input
		aStr, op, bStr := parts[0], parts[1], parts[2]

		// Konversi ke float64
		a, err1 := strconv.ParseFloat(aStr, 64)
		b, err2 := strconv.ParseFloat(bStr, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Input angka tidak valid!")
			continue
		}

		// Hitung sesuai operator
		switch op {
		case "+":
			fmt.Printf("Hasil: %.6g %s %.6g = %.6g\n", a, op, b, a+b)
		case "-":
			fmt.Printf("Hasil: %.6g %s %.6g = %.6g\n", a, op, b, a-b)
		case "*":
			fmt.Printf("Hasil: %.6g %s %.6g = %.6g\n", a, op, b, a*b)
		case "/":
			if b == 0 {
				fmt.Println("Error: pembagian dengan nol")
				continue
			}
			fmt.Printf("Hasil: %.6g %s %.6g = %.6g\n", a, op, b, a/b)
		case "%":
			if b == 0 {
				fmt.Println("Error: modulo dengan nol")
				continue
			}
			fmt.Printf("Hasil: %.6g %s %.6g = %.6g\n", a, op, b, math.Mod(a, b))
		default:
			fmt.Println("Operator tidak dikenal. Gunakan + - * / %")
		}
	}
}
