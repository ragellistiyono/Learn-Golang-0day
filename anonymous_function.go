package main

import "fmt"

type funcChecking func(string) string

func halamanAdmin(username string, checking funcChecking) string {
	fmt.Println("==== Tahapan Checking username =====")
	fmt.Print("Silahkan Masukkan username : ")
	fmt.Scanf("%s", &username)
	return checking(username)
}

func main() {
	// Membuat anonymous function
	checkUsername := func(username string) string {
		if username == "ragel" {
			return "Selamat Datang"
		} else {
			return "Kamu tidak lolos tahapan checking"
		}
	}

	// Memanggil fungsi halamanAdmin dengan anonymous function
	result := halamanAdmin("", checkUsername)
	fmt.Println(result)
}
