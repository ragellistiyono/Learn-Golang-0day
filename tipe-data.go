package main

import "fmt"

func main() {
	gel := `
  ___               ___  
 (o o)             (o o) 
(  V  ) humblegod (  V  )
--m-m---------------m-m--
`
	bilanganPositif := 1234
	bilanganNegatif := -1234
	bilanganDesimal := 129.27
	benar := true  //tipe data boolean
	salah := false //tipe data boolean

	fmt.Println(gel)
	fmt.Printf("Bilangan Positif %d \n", bilanganPositif)
	fmt.Printf("Bilangan Negatif %d\n", bilanganNegatif)
	fmt.Printf("Bilangan Negatif %.2f\n", bilanganDesimal)
	fmt.Printf("Benar %t\n", benar)
	fmt.Printf("Salah %t\n", salah)

}
