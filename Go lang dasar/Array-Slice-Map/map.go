package main

import "fmt"

func main() {
	// membuat map
	person := map[string]string{
		"nama":   "Wiranto",
		"addres": "subang",
	}
	fmt.Println(person["nama"])
	fmt.Println(len(person))

	// function pada map
	// len(map)
	// map["key"]
	// map["key"] = value
	// make(map[int]int)

	personBaru := make(map[string]string)
	personBaru["nama"] = "vicry"
	personBaru["address"] = "ROCK_BOTTOM"

	fmt.Println(personBaru["nama"])
	fmt.Println(personBaru["address"])
}
