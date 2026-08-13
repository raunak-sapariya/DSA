/*
CIRCULAR QUEUE
	two pointers front and rear
		initially both front and rear are set to -1
		front points to the first element of the queue
		rear points to the last element of the queue

		if inserting first element, both front and rear are set to 0
		if inserting subsequent elements, rear is incremented by 1
		if deleting an element, front is incremented by 1

		(rare + 1) % size -> if rear is at the end of the array and there is space at the beginning of the array,
						we can wrap around and insert the element at the beginning of the array

		full -> if (rear + 1) % size == front
		empty -> if front == -1
*/

package main

import "fmt"

func enqueue(queue []int, size *int, front *int, rare *int, value int) []int {
	if (*rare+1)%*size == *front {
		fmt.Println("Queue is full")
		return queue
	}
	*rare = (*rare + 1) % *size
	queue[*rare] = value
	if *front == -1 {
		*front = 0
	}
	return queue
}

func dequeue(queue []int, size *int, front *int, rare *int) {
	if *front == -1  {
		fmt.Println("Queue is empty")
		return
	}
	value := queue[*front]
	*front = (*front + 1) % *size
	if *front == *rare {
		*front = -1
		*rare = -1
	}
	fmt.Println("Dequeued value: ", value)
}

func display(queue []int, size *int, front *int, rare *int) {
	if *front == -1 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Print("Queue elements: ")
	for i := *front; i != *rare; i = (i + 1) % *size {
		fmt.Print(queue[i], " ")
	}
	fmt.Println(queue[*rare])
}

func main() {
	size := 5
	front := -1
	rare := -1
	queue := make([]int, size)
	fmt.Println("Initial Queue: ", queue)
	
	queue = enqueue(queue, &size, &front, &rare, 10)
	queue = enqueue(queue, &size, &front, &rare, 20)
	queue = enqueue(queue, &size, &front, &rare, 30)
	queue = enqueue(queue, &size, &front, &rare, 40)
	queue = enqueue(queue, &size, &front, &rare, 50)
	display(queue, &size, &front, &rare)

	// Full
	queue = enqueue(queue, &size, &front, &rare, 50)

	dequeue(queue, &size, &front, &rare)
	dequeue(queue, &size, &front, &rare)
	dequeue(queue, &size, &front, &rare)
	dequeue(queue, &size, &front, &rare)
	dequeue(queue, &size, &front, &rare)
	// empty
	dequeue(queue, &size, &front, &rare)
	dequeue(queue, &size, &front, &rare)
	display(queue, &size, &front, &rare)
}