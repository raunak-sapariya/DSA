package main

import "fmt"

func displayMatrix(arr [][]int, rows, cols int) {
	fmt.Println("\nArray after rotation:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var rows, cols int

	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter number of columns: ")
	fmt.Scan(&cols)

	arr := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
	}

	fmt.Println("\nEnter array elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			fmt.Scan(&arr[i][j])
		}
	}

	fmt.Println("\nOriginal array:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

	var count int
	fmt.Print("\nEnter rotation count: ")
	fmt.Scan(&count)

	// If count is greater than columns
	count = count % cols

	for i := 0; i < rows; i++ {

		// Rotate current row 'count' times
		for k := 0; k < count; k++ {

			// last element of current row
			last := arr[i][cols-1]

			// shifting all elements of current row by 1
			for j := cols - 1; j > 0; j-- {
				arr[i][j] = arr[i][j-1]
			}

			// Place last element at first position
			arr[i][0] = last
		}
	}

	displayMatrix(arr, rows, cols)
}