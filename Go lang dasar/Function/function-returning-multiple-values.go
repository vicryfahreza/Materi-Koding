package main

import "fmt"

// function tak hanya dapat mengembalikan satu value, tapi juga bisa multiple value.
// Untuk memberitahu jika function mengembalikan multiple value,Kita harus menulis semua tipe data return valuenya di function.
func getFullName() (string, string) {
	return "Vicry", "Fahreza"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)

}
