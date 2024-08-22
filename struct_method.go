package main

import "fmt"

type Badminton struct {
	Name    string
	Country string
	Age     int
}

func (badminton Badminton) PlayerName() string {
	return badminton.Name
}

func (badminton Badminton) PlayerCountry() string {
	return badminton.Country
}

func (badminton Badminton) PlayerAge() int {
	return badminton.Age
}

func main() {
	axelsen := Badminton{
		Name:    "Victor Axelsen",
		Country: "Denmark",
		Age:     30,
	}
	fmt.Println(axelsen.PlayerName())
	fmt.Println(axelsen.PlayerCountry())
	fmt.Println(axelsen.PlayerAge())
}
