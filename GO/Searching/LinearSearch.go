/*
Time complexity: O(n)
*/
package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var target int
	fmt.Print("Enter the element to search: ")
	fmt.Scan(&target)
	result := LinearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found in the array.")
	}
}