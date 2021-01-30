package main

import "fmt"

func main() {
	// hanya ada for loop di go-lang
	for i := 0; i <= 5; i++ {
		fmt.Println("loop ke-", i)
	}

	// for range array dan slice
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index : ", index, "value : ", value)
	}

	// for range map
	alphabet := make(map[string]string)
	alphabet["a"] = "alpha"
	alphabet["b"] = "beta"

	for key, value := range alphabet {
		fmt.Println("key : ", key, "value : ", value)
	}

}
