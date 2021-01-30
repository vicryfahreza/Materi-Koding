package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(12, 12, 12)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	total = sumAll(slice...)
	fmt.Println(total)

}
