package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	const distance = 62100000

	n, _ := fmt.Printf("%-20v %5v %10v %5v\n", "Spaceline", "Days", "Round-trip", "Price")
	line := strings.Repeat("=", n)
	fmt.Println(line)

	for i := 0; i < 10; i++ {
		spaceline := rand.Intn(3) + 1
		var spacelineName string

		switch spaceline {
		case 1:
			spacelineName = "Space Adventures"
		case 2:
			spacelineName = "SpaceX"
		case 3:
			spacelineName = "Virgin Galactic"
		}

		fmt.Printf("%-20s ", spacelineName)

		priceRaise := rand.Intn(15)
		speed := priceRaise + 16
		fmt.Printf("%5v ", distance/speed/60/60/24)

		roundTrip := rand.Intn(2) == 1

		if roundTrip {
			fmt.Printf("%-10v ", "Round-trip")
			fmt.Printf("$%4v \n", (priceRaise+36)*2)
		} else {
			fmt.Printf("%-10v ", "One-way")
			fmt.Printf("$%4v \n", priceRaise+36)
		}
	}
}
