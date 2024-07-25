package main

import "fmt"

func main() {

	badminton := []string{"Ginting", "Momota", "Axelsen", "Jonathan"}
	jagoan := badminton[0:4]

	fmt.Println("Player Badminton", badminton)
	fmt.Println("Number 1", badminton[0])
	fmt.Println("Number 2", badminton[1])
	fmt.Println("Number 3", badminton[2])
	fmt.Println("Number 4", badminton[3])
	fmt.Println(jagoan)
	fmt.Println()

	aJagoan := badminton[0:3]
	bJagoan := badminton[1:4]

	aaJagoan := aJagoan[1:2]
	bbJagoan := bJagoan[0:1]

	fmt.Println(aJagoan)
	fmt.Println(bJagoan)
	fmt.Println(aaJagoan)
	fmt.Println(bbJagoan)
	fmt.Println("Pendatang baru Badminton")

	bbJagoan[0] = "Ragel"
	fmt.Println(aJagoan)
	fmt.Println(bJagoan)
	fmt.Println(aaJagoan)
	fmt.Println(bbJagoan)

}
