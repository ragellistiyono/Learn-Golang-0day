package main

import "fmt"

func main() {
	//Deklarasi Variable manifest typing
	var firstName string = "Ragel"
	var lastName string = "Listiyono"

	fmt.Println("Hallo, Saya", firstName, lastName)

	//Deklarasi Variable tanpa tipe data
	asalKota := "Surabaya"
	usia := "26"

	fmt.Printf("Saya Berasal dari %s dan berumur %s ", asalKota, usia)
	fmt.Println()

	//Deklarasi multi Variabel

	hobby, status, inspired := "Badminton", "Mahasiswa", "Lionel Messi"

	fmt.Printf("Hobby saya %s, status saya saat ini adalah %s, inspirasi saya %s", hobby, status, inspired)
}
