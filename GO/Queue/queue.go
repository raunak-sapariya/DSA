/*
QUEUE
	First in first out
	enqueue() - O(1) insert at rear
	dequeue() - O(1) remove from front
	front()   - O(1) access first element
	rear()    - O(1) access last element
	size()    - O(1) get number of elements
	display() - O(n) display all elements
*/
package main

import "fmt"

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
		// Dequeue the element at the front of the queue
		if *front <= *rare {
			fmt.Println("\nDequeued element:", queue[*front])
			*front = *front + 1
			// Reset front and rare if the queue becomes empty after dequeue
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
		fmt.Print("\nQueue elements: ")
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

	queue = enqueue(queue[:], size, &front, &rare, 10)
	queue = enqueue(queue[:], size, &front, &rare, 20)
	queue = enqueue(queue[:], size, &front, &rare, 30)
	fmt.Println("\nQueue after enqueuing 10, 20, 30:")
	display(queue[:], &front, &rare)

	queue = dequeue(queue[:], &front, &rare)
	fmt.Println("Queue after dequeuing an element:")
	display(queue[:], &front, &rare)
}