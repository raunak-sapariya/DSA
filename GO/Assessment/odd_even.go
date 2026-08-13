package main

import "fmt"

func main() {

		var n int

		fmt.Print("Enter an Number of integer: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Error: Invalid input! Please enter an integer.")
		}

		even := make([]int,0)
		odd := make([]int,0)

		for i := 0; i < n; i++{
			var n int
			fmt.Printf("Enter a number %v: ",i+1)
			_, err = fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error: Invalid input! Please enter an integer.")
				i --
			}
			

			if n%2 == 0 {
				even = append(even, n)
			} else {
				odd = append(odd, n)
			}
		}
		
		fmt.Println("Even",even)
		fmt.Println("Odd",odd)
}
