package main

import "fmt"

func namaAtlit() {
	fmt.Println("Ragel Listiyono")
}

func main() {

	defer namaAtlit()

	//jika terjadi error defer ttp tereksekusi asalkan di tempatkan di sebelum code error/ atas

	ginting := 0
	result := 10 / ginting

	fmt.Println(result)

	fmt.Println("Anthony Ginting")
	fmt.Println("Victor Axelsen")

}
