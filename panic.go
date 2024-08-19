package main

import "fmt"

func aritmatika() {
	var angka1, angka2 int

	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&angka1)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scan(&angka2)

	if angka2 == 0 {
		panic("Tidak dapat membagi dengan nol")
	}
	hasil := angka1 / angka2

	fmt.Printf("Hasil dari %d dibagi %d adalah %d\n", angka1, angka2, hasil)

}

func main() {
	aritmatika()
}
