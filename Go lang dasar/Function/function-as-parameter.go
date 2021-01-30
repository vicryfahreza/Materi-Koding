package main

import "fmt"

type typeFilter func(string) string

func sayHelloWithFilter(name string, filter typeFilter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Vicry", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
