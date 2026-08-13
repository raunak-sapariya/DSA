package main

import (
	"fmt"
)


func main() {
	var n int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nOriginal array:", arr)

	// Insert
	var position, value int
	fmt.Printf("\nEnter the position to insert the element (0 to %d): ", len(arr)-1)
	fmt.Scan(&position)

	if position < 0 || position >= len(arr) {
		fmt.Println("Invalid position!")
		return
	}

	fmt.Print("Enter the value to insert: ")
	fmt.Scan(&value)

	arr = append(arr[:position], append([]int{value}, arr[position:]...)...)	

	fmt.Println("Array after insertion:", arr)

	// Delete
	fmt.Printf("\nEnter the position to delete the element (0 to %d): ", len(arr)-1)
	fmt.Scan(&position)

	if position < 0 || position >= len(arr) {
		fmt.Println("Invalid position!")
		return
	}

	arr = append(arr[:position], arr[position+1:]...)
	fmt.Println("Array after deletion:", arr)
}