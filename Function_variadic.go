package main

import "fmt"

func penjumlahan(angka ...int) int {
	total := 0

	for _, nomor := range angka {
		total += nomor
	}
	return total
}

func main() {
	hasil := (penjumlahan(5, 5, 5, 5))
	fmt.Println(hasil)
}
