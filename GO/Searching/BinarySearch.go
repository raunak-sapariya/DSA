/*
Time complexity: O(log n)
*/
package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	// arr := []int{10, 20, 30, 40, 50, 60}
	var arr []int
	var n int

	fmt.Print("wEnter the number of elements in the array: ")
	fmt.Scan(&n)

	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Enter element ", i+1, ": ")
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array",arr)

	var target int
 	fmt.Print("Enter the element to search: ")
	fmt.Scan(&target)

	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found in the array.")
	}
}