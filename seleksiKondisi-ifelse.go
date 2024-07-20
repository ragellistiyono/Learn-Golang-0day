package main

import "fmt"

func main() {
	nilai := 65

	fmt.Printf("Masukkan Nilai Ujian : ")
	fmt.Scanf("%d", &nilai)

	if nilai == 100 {
		fmt.Printf("Selamat Nilai Sempurna! Kamu tidak mengikuti Remidi\n")
	} else if nilai >= 65 {
		fmt.Printf("Selamat Kamu tidak mengikuti Remidi")
	} else {
		fmt.Printf("Kamu Harus mengikuti Remidi")
	}
}
