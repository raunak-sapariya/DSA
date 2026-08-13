// front - deletion (dequeue)
// rare - insertion (enqueue)

package main

import (
	"fmt"
)

func enqueue(queue []int, size int, front *int, rare *int, value int) []int {
	if *rare == -1 {
		*front = 0
		*rare = 0
		queue = append(queue, value)
	} else if *rare == size-1 {
		fmt.Println("Queue is full. Cannot enqueue.")
	} else {
		*rare = *rare + 1
		queue = append(queue, value)
	}
	return queue
}

func dequeue(queue []int, front *int, rare *int) []int {
	if *front == -1 {
		fmt.Println("Queue is empty. Cannot dequeue.")
	} else {
		if *front <= *rare {
			fmt.Println("Dequeued element:", queue[*front])
			*front = *front + 1
			if *front > *rare {
				*front = -1
				*rare = -1
			}
		}
	}
	return queue
}

func display(queue []int, front *int, rare *int) {
	if *front == -1 {
		fmt.Println("Queue is empty.")
	} else {
		fmt.Print("Queue elements: ")
		for i := *front; i <= *rare && i < len(queue); i++ {
			fmt.Print(queue[i], " ")
		}
		fmt.Println()
	}
}

func main(){
	size := 5
	var queue []int
	var front, rare = -1, -1

	fmt.Println("Initial Queue:")
	display(queue[:], &front, &rare)

	fmt.Print("\nEnter number of elements to enqueue: ")
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter element ", i+1, ": ")
		var value int
		fmt.Scan(&value)
		queue = enqueue(queue[:], size, &front, &rare, value)
	}
	fmt.Println(queue)
	display(queue[:], &front, &rare)

	fmt.Print("\nEnter number of elements to dequeue: ")
	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		queue = dequeue(queue[:], &front, &rare)
	}
	fmt.Println(queue)
	display(queue[:], &front, &rare)
}