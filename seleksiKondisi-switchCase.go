package main

import "fmt"

func main() {

	minimNilaiUjian := 50

	fmt.Printf("Masukkan Nilai :")
	fmt.Scanf("%d", &minimNilaiUjian)

	switch {
	case minimNilaiUjian == 100:
		fmt.Printf("Selamat Nilai Kamu sempurna")
	case (minimNilaiUjian < 100) && (minimNilaiUjian > 50):
		fmt.Printf("Kamu tidak Remidi")
	default:
		{
			fmt.Printf("Kamu harus Remidi\n")
			fmt.Printf("Belajar lebih giat")
		}
	}
}
