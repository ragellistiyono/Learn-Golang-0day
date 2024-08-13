package main

import "fmt"

func getSelamatDatang(nama string) string {
	return "Selamat Datang " + nama
}

func main() {

	hasil := getSelamatDatang

	fmt.Println(hasil("Ragel Listiyono"))
}
