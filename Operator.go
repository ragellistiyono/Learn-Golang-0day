package main

import "fmt"

func main() {
	//Operator Aritmatika
	hitungan1 := (((2+3)*2)*8 + 2) / 2

	fmt.Printf("hasilnya adalah %d\n", hitungan1)
	fmt.Printf("===============================\n")

	//Operator Perbandingan
	angka := 10*10 + 10
	hasilAngka := (angka == 120)
	hasilAngka1 := (angka < 120)
	hasilAngka2 := (angka > 120)

	fmt.Printf("Operator perbandingan dari 10*10+10\n")
	fmt.Printf("Apakah hasilnya sama dengan %d (%t)\n", angka, hasilAngka)
	fmt.Printf("Apakah hasil lebih besar dari %d (%t)\n", angka, hasilAngka2)
	fmt.Printf("Apakah hasil kurang dari %d (%t)\n", angka, hasilAngka1)
}
