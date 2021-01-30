package main

import "fmt"

// Membuat variable secara langsung di tipe data return function nya
func returnVar() (firstName string, middleName string, lastName string) {
	firstName = "Vicry"
	middleName = "justify"
	lastName = "Fahreza"
	return
}

func main() {
	a, b, c := returnVar()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
