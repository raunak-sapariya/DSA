
/*
STACK
	First in last out
	push() 	  - O(1) insert at top
	pop()  	  - O(1) remove from top
	peek() 	  - O(1) access top element
	size() 	  - O(1) get number of elements
	display() - O(n) display all elements
*/
package main

import (
	"fmt"
)

func push(stack []int, size int, top *int, value int) []int {
	if *top == size-1 {
		fmt.Println("Stack is full. Cannot push.")
	} else {
		*top = *top + 1
		stack = append(stack, value)
	}
	return stack
}

func pop(stack []int, top *int) []int {
	if *top == -1 {
		fmt.Println("\nStack is empty. Cannot pop.")
	} else {
		fmt.Println("\nPopped element:", stack[*top])
		stack = stack[:*top] // Remove the top element
		*top = *top - 1
	}
	return stack
}

func peek(stack []int, top *int) {
	if *top == -1 {
		fmt.Println("Stack is empty. No top element.")
	} else {
		fmt.Println("Top element:", stack[*top])
	}
}

func size(stack []int, top *int) {
	fmt.Println("Current size of stack:", *top+1)
}

func display(stack []int, top *int) {
	if *top == -1 {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Print("Stack elements: ")
		for i := 0; i <= *top; i++ {
			fmt.Print(stack[i], " ")
		}
		fmt.Println()
	}
}

func main() {
	stack_size := 5
	var stack []int
	var first = -1

	fmt.Println("Initial Stack:")
	display(stack[:], &first)

	stack = push(stack, stack_size, &first, 10)
	stack = push(stack, stack_size, &first, 20)
	stack = push(stack, stack_size, &first, 30)

	fmt.Println("\nStack after pushing elements 10, 20, 30:")
	display(stack[:], &first)

 	pop(stack[:], &first)
	size(stack[:], &first)
	display(stack[:], &first)

	pop(stack, &first)
	display(stack[:], &first)

	pop(stack, &first)
	display(stack[:], &first)

	fmt.Println("\nStack after popping elements:")
	display(stack[:], &first)
 	peek(stack[:], &first)
	size(stack[:], &first)
}