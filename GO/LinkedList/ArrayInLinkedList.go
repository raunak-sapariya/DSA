package main

import "fmt"


type Node struct {
	data []int
	size int
	next *Node
}

type ArrayInLinkedList struct {
	head    *Node
	tail    *Node
}

func (all *ArrayInLinkedList) InsertAtEnd(data []int, size int) {
	newNode := &Node{data: data, size: size}
	if all.head == nil {
		all.head = newNode
		all.tail = newNode
	} else {
		all.tail.next = newNode
		all.tail = newNode
	}
}

func (lll *ArrayInLinkedList) display(){
	if lll.head == nil {
		fmt.Println("List is empty.")
	} else {
		current := lll.head
		for current != nil {
			fmt.Print(current.data, "Size: ", current.size)
			if (current.next != nil){
				fmt.Print("->")
			}
			current = current.next
		}
	}
}


func main() {
	all := &ArrayInLinkedList{}
	var linkedListSize int
	fmt.Print("Enter the size of Linked List: ")
	fmt.Scan(&linkedListSize)

	for i := 0; i < linkedListSize; i++ {
		var n int
		fmt.Printf("Enter the size of array for node %d: ", i+1)
		fmt.Scan(&n)

		data := make([]int, n)
		fmt.Printf("Enter the elements of array %d: ", i+1)
		for j := 0; j < n; j++ {
			fmt.Scan(&data[j])
		}
		all.InsertAtEnd(data, n)
	}
	fmt.Println()
	all.display()
}