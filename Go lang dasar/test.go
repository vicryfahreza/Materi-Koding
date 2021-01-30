package main

import "fmt"

func main() {
	for a := 0; a <= 10; a++ {
		for b := 0; b < a; b++ {
			fmt.Print("$ ")
		}
		fmt.Println()
	}
	fmt.Println()
	for c := 10; c >= 1; c-- {
		for d := 0; d < c; d++ {
			fmt.Print("$ ")
		}
		fmt.Println()
	}
}
