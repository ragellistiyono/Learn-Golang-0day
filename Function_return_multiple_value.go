package main

import "fmt"

func sport() (string, string, int) {
	return "Saya adalah " + "Ragel", "Player " + "Badminton", 10
}

func main() {
	namaDepan, Olahraga, nomor := sport()
	fmt.Println(namaDepan, Olahraga, nomor)

	namaDepan, Olahraga, _ = sport()
	fmt.Println(namaDepan, Olahraga)
}
