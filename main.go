package main

import (
	"fmt"
	"strings"
)

func main() {
	n, _ := fmt.Printf("%-15v %-4v %-10v %-5v\n", "Spaceline", "Days", "Round-trip", "Price")
	line := strings.Repeat("=", n)
	fmt.Println(line)
}