package main

import "fmt"

func main() {
	fmt.Println("masukkan nomor:")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println("hasil:", output)
}
