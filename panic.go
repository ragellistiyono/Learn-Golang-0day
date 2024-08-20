package main

import "fmt"

func aritmatika() {

	defer mengatasiPanic()
	var angka1, angka2 float64

	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&angka1)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scan(&angka2)

	if angka2 == 0 {
		panic("Tidak dapat membagi dengan nol")
	}
	hasil := angka1 / angka2

	fmt.Printf("Hasil dari %.1f dibagi %.1f adalah %.1f\n", angka1, angka2, hasil)

}

func mengatasiPanic() {
	err := recover()

	if err != nil {
		fmt.Println("Ini adalah RECOVER agar pesan error tidak tampil | ", err)
	}
}

func main() {
	aritmatika()
}
