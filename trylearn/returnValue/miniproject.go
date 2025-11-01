package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Masukkan nilai dan tekan Enter (pisahkan dengan spasi):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim & split (mengabaikan spasi berganda)
	check := strings.TrimSpace(input)
	if check == "" {
		fmt.Println("Input tidak boleh kosong!")
		return
	}
	parts := strings.Fields(check)

	// Parse semua elemen menjadi slice of int
	nums := make([]int, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Input '%s' bukan angka valid. Harap masukkan hanya angka dan spasi.\n", s)
			return
		}
		nums = append(nums, n)
	}

	// Sekarang hitung rata-rata dari seluruh nums
	rata := rataRata(nums...)
	fmt.Printf("Rata-rata: %.2f\n", rata)

	// Inisialisasi max/min dari elemen pertama (karena nums tidak kosong)
	nilaiTertinggi := nums[0]
	nilaiTerendah := nums[0]
	var jumlahDiatasRata int

	// Loop untuk cari max/min dan hitung berapa yang > rata
	for _, n := range nums {
		if n > nilaiTertinggi {
			nilaiTertinggi = n
		}
		if n < nilaiTerendah {
			nilaiTerendah = n
		}
		if float64(n) > rata {
			jumlahDiatasRata++
		}
	}

	fmt.Printf("Nilai tertinggi: %d\n", nilaiTertinggi)
	fmt.Printf("Nilai terendah: %d\n", nilaiTerendah)
	fmt.Printf("Jumlah nilai di atas rata-rata: %d\n", jumlahDiatasRata)
}

// rataRata menerima variadic int dan mengembalikan float64
func rataRata(angka ...int) float64 {
	if len(angka) == 0 {
		return 0
	}
	total := 0
	for _, v := range angka {
		total += v
	}
	return float64(total) / float64(len(angka))
}
