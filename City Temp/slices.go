package main

import "fmt"

func main3() {
	array := [3]string{"a", "b", "c"}

	copy1 := array[0:2]
	copy2 := array[1:3]
	// creates a new independent slice
	copy3 := append(copy2, "flan")

	array[0] = "x"
	updateSlice(copy1, "y")
	updateSlice(copy2, "z")

	fmt.Println("Array: ", array)
	fmt.Println("Copy1: ", copy1)
	fmt.Println("Copy2: ", copy2)
	fmt.Println("Copy3: ", copy3)
}

// slice is a pointer to an array under the hood, so no need for explicit *
func updateSlice(slice []string, value string) {
	slice[1] = value
}
