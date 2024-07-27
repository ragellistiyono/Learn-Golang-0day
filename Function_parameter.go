package main

import "fmt"

func playerBadminton(namaDepan string, ranking int) {
	fmt.Println("Halo", namaDepan, "Player Ranking", ranking)
}

func main() {
	playerBadminton("Axelsen", 1)
	playerBadminton("Ginting", 2)
	playerBadminton("Jonathan", 3)
}
