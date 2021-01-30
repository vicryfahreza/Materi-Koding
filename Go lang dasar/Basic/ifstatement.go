package main

import "fmt"

func main() {
	// if short statement
	name := "JamesWreck"

	// mudahnya
	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("ok pass")
	}

	// panjangnya
	length := len(name)
	if length > 5 {
		fmt.Println("terlalu panjang ifnya")
	} else {
		fmt.Println("ok pass")
	}
}
