package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) InsertAtBegin(data int) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoublyLinkedList) DeleteFromEnd() {
	if dll.head == nil {
		fmt.Println("List is empty.")
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
}

func (dll *DoublyLinkedList) DeleteFromBegin() {
	if dll.head == nil {
		fmt.Println("List is empty.")
		return
	}

	// Only one node
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return
	}

	dll.head = dll.head.next
	dll.head.prev = nil
}

func (dll *DoublyLinkedList) Display() {
	if dll.head == nil {
		fmt.Println("Empty")
		return
	}

	current := dll.head
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" <-> ")
		}
		current = current.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) DisplayReverse() {
	if dll.tail == nil {
		fmt.Println("Empty")
		return
	}

	current := dll.tail
	for current != nil {
		fmt.Print(current.data)
		if current.prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.prev
	}
	fmt.Println("")
}

func main() {
	dll := &DoublyLinkedList{}

	dll.InsertAtEnd(10)
	dll.InsertAtEnd(20)
	dll.InsertAtEnd(30)

	fmt.Println("After InsertAtEnd:")
	dll.Display()

	dll.InsertAtBegin(5)
	fmt.Println("After InsertAtBegin:")
	dll.Display()

	dll.DeleteFromEnd()
	fmt.Println("After DeleteFromEnd:")
	dll.Display()

	dll.DeleteFromBegin()
	fmt.Println("After DeleteFromBegin:")
	dll.Display()

	fmt.Println("Reverse Traversal:")
	dll.DisplayReverse()
}