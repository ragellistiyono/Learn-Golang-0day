package main

import "fmt"

func main() {
	//Array Vertikal
	var nama = [4]string{"ragel", "messi", "ralisto", "reddington"}

	//Array Horizontal
	var olahraga = [2]string{
		"bola",
		"badminton",
	}

	//Array dengan gaya type inference dan tipe data campur menggunakan inteface
	campuran := [4]interface{}{"Ragel", 10, "Listiyono", 19.27}

	//Array tanpa jumlah elemen
	angka := [...]int{1, 2, 3, 4}

	//Array Multidimensi
	multiDimensi := [2][4]string{
		{
			"Ragel",
			"Listiyono",
			"10",
			"September"},
		{
			"Surabaya",
			"Indonesia",
			"Laki-Laki",
			"Mahasiswa"}}

	// Array for
	buah := [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Nama :\t\t", nama)
	fmt.Println("Olahraga:\t", olahraga)
	fmt.Println("Data Campur:\t", campuran)
	fmt.Println("Angka: \t\t", angka)
	fmt.Println("Array Multi:\t", multiDimensi)
	for i, buah := range buah {
		fmt.Printf("elemen %d: %s\n", i, buah)
	}
}
