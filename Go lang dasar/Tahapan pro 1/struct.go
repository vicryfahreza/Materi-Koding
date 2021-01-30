package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// struct mirip objek
	var vic Customer
	vic.Name = "vicy"
	vic.Address = "jln suka bulan"
	vic.Age = 18

	fmt.Println(vic.Name)
	fmt.Println(vic.Address)
	fmt.Println(vic.Age)

	// struct literal
	Wanto := Customer{
		Name:    "Wanto pede",
		Address: "jln mengucap salah",
		Age:     60,
	}
	fmt.Println(Wanto.Name)

	Budi := Customer{"Budi gemuk ", "jln menghadap suridman", 30}
	fmt.Println(Budi)
}
