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

	count = count % cols

	for j := 0; j < cols; j++ {

		for k := 0; k < count; k++{

			last := arr[rows-1][j]
			
			for i := rows - 1; i > 0; i--{
				arr[i][j] = arr[i-1][j]
			}

			arr[0][j] = last
		}
		
	}


	displayMatrix(arr, rows, cols)
}