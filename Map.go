package main

import "fmt"

func main() {
	player := map[string]int{}
	player["Ginting"] = 3
	player["Jonathan"] = 6

	//map didefinisikan di awal dengan gaya penulisan vertikal
	olahraga := map[string]int{
		"basket": 1,
		"voli":   2,
	}

	// map dengan for-range
	ranking := map[int]string{
		1: "Axelsen",
		2: "Ginting",
		3: "Jonathan",
		4: "Antonsen",
		5: "Momota",
	}

	//Kombinasi Map dan SLice
	sepakBola := []map[string]string{
		{"nama": "Lionel Messi", "club": "Barcelona"},
		{"nama": "Christiano Ronaldo", "club": "Real Madrid"},
	}

	fmt.Println("Ranking Player Ginting :", player["Ginting"])
	fmt.Println("Ranking Player Jonathan:", player["Jonathan"])
	fmt.Println(olahraga)
	fmt.Println("Ranking PLayer Badminton")
	for key, value := range ranking {
		fmt.Println(key, " \t", value)
	}
	for _, bola := range sepakBola {
		fmt.Println(bola["nama"], bola["club"])
	}
}
