package main

import "fmt"

func badminton() (namaDepan, namaBelakang, negara string, ranking int) {
	namaDepan = "Ragel"
	namaBelakang = "Listiyono"
	negara = "Indonesia"
	ranking = 1

	return "hallo " + namaDepan, namaBelakang, negara, ranking
}

func main() {
	namaDepan, namaBelakang, negara, ranking := badminton()
	fmt.Println(namaDepan, namaBelakang, negara, ranking)
}
