package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	//var distance = 62100000

	n, _ := fmt.Printf("%-20v %-5v %-10v %-5v\n", "Spaceline", "Days", "Round-trip", "Price")
	line := strings.Repeat("=", n)
	fmt.Println(line)

	for i := 0; i < 10; i++ {
		airline := rand.Intn(3) + 1

		switch airline {
		case 1:
			fmt.Printf("%-20v\n", "Space Adventures")
		case 2:
			fmt.Printf("%-20v\n", "SpaceX")
		case 3:
			fmt.Printf("%-20v\n", "Virgin Galactic")
		}
	}
}
