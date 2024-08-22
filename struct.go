package main

import "fmt"

type Badminton struct {
	Nama    string
	Address string
	Umur    int
}

func main() {
	var player Badminton
	player.Nama = "Ragel Listiyono"
	player.Address = "Surabaya"
	player.Umur = 26

	fmt.Println(player)

	//Struct Literals

	ginting := Badminton{
		Nama:    "Anthony Ginting",
		Address: "Indonesia",
		Umur:    25,
	}
	fmt.Println(ginting)
}
