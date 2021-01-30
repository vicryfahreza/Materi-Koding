package main

import (
	"fmt"
)

func main() {
	fmt.Println("mengenal variable pada golang")

	var name string

	name = "namaku beno aku suka bento"
	fmt.Println("no.1 ", name)

	name = "namaku joni aku suka ngoli"
	fmt.Println("no.2 ", name)

	var mantanname = "jesica cleopatra"
	fmt.Println(mantanname)

	var age = 18
	fmt.Println(age)

	// var tidak wajib
	namaku := "totong"

	fmt.Println("aku ", namaku)

	var (
		firstname = "vicry"
		lastname  = "fahreza"
	)
	fmt.Println(firstname, lastname)

}
