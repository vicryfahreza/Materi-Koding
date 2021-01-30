package main

import "fmt"

func fullname() (string, string) {
	return "vicry", "fahreza"
}

func main() {
	a, b := fullname()
	fmt.Println(a)
	fmt.Println(b)

}
