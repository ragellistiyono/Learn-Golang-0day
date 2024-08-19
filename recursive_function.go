package main

import "fmt"

func perulanganFaktorial(nominal int) int {
	if nominal == 1 {
		return 1
	} else {
		return nominal * perulanganFaktorial(nominal-1)
	}

}

func main() {

	bilangan := 0
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scan(&bilangan)

	hasil := perulanganFaktorial(bilangan)
	fmt.Printf("faktorial dari %d adalah %d\n", bilangan, hasil)

}
