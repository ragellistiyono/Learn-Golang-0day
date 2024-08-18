package main

import "fmt"

func filterPesan(kata string, filter func(string) string) {
	filteredPesan := filter(kata)
	fmt.Println("Selamat Datang", filteredPesan)
}

func kataJorok(kata string) string {
	fmt.Printf("Masukkan Kata : ")
	fmt.Scanf("%s", &kata)
	if kata == "Babi" {
		return "Kata yang kamu input tidak diperbolehkan"
	} else {
		return kata
	}
}
func main() {

	filter := kataJorok
	filterPesan("Babi", filter)
}
