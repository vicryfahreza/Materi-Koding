package main

import "fmt"

func main() {
	months := [...]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(months)
	slice1 := months[4:7]
	// function slice len() dan cap()
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println(slice1)
	slice1[0] = "mei_lagi"
	slice1[1] = "juni_lagi"

	fmt.Println(slice1)
	fmt.Println(months)

	// function slice append()
	Hari := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	fmt.Println(Hari)
	slice_a := Hari[4:]
	slice_a[0] = "jumat_berkah"
	slice_a[1] = "sabtu_mantap"
	fmt.Println(slice_a)

	slice_b := append(slice_a, "minggu_baru")
	slice_b[2] = "ini_minggu"
	fmt.Println(slice_b)

	// function make
	slice_make := make([]string, 3, 5)
	slice_make[0] = "vicry"
	slice_make[1] = "fahreza"
	slice_make[2] = "marinir"
	fmt.Println(slice_make)
	fmt.Println(len(slice_make))
	fmt.Println(cap(slice_make))

	// function copy
	slice_copy := make([]string, 2, cap(slice_make))
	copy(slice_copy, slice_make)
	fmt.Println(slice_copy)

	// waspada slice itu sama seperti array
	// Slice_x := []string{}

	// Array_x := [...]string{}

}
