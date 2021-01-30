package main

import "fmt"

func functionOnVar(nama string) string {
	return "Hallo " + nama
}

func main() {
	Variabel := functionOnVar
	result := Variabel("Vicry")
	fmt.Println(result)
}
