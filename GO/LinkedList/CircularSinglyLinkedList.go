package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularSinglyLinkedList struct {
	head *Node
	tail *Node
}

func (csll *CircularSinglyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if csll.head == nil {
		csll.head = newNode
		csll.tail = newNode
		csll.tail.next = csll.head
	} else {
		csll.tail.next = newNode
		csll.tail = newNode
		csll.tail.next = csll.head
	}
}

func (csll *CircularSinglyLinkedList) InsertAtBegin(data int) {
	newNode := &Node{data: data}
	if csll.head == nil {
		csll.head = newNode
		csll.tail = csll.head
		csll.tail.next = csll.head
	} else {
		newNode.next = csll.head
		csll.head = newNode
		csll.tail.next = csll.head
	}
}

func (csll *CircularSinglyLinkedList) DeleteFromEnd() {
	if csll.head == nil {
		return
	}

	// one node
	if csll.head == csll.tail {
		csll.head = nil
		csll.tail = nil
		return
	}

	current := csll.head
	for current.next != csll.tail {
		current = current.next
	}
	current.next = csll.head
	csll.tail = current
}

func (csll *CircularSinglyLinkedList) DeleteFromBegin() {
	if csll.head == nil {
		return
	}

	// one node
	if csll.head == csll.tail {
		csll.head = nil
		csll.tail = nil
		return
	}

	csll.head = csll.head.next
	csll.tail.next = csll.head

}

func (csll *CircularSinglyLinkedList) DeleteTail() {
	if csll.head == nil {
		fmt.Println("Circular Singly Linked List is empty. Cannot delete from tail.")
		return
	}

	// one node
	if csll.head == csll.tail {
		csll.head = nil
		csll.tail = nil
		return
	}

	temp := csll.head
	for temp.next != csll.tail {
		temp = temp.next
	}
	temp.next = csll.head
	csll.tail = temp
}

func (csll *CircularSinglyLinkedList) Display() {
	if csll.head == nil {
		fmt.Println("Circular Singly Linked List is empty.")
		return
	} else {
		current := csll.head
		for {
			fmt.Print(current.data, "->")
			current = current.next
			if current == csll.head {
				break
			}
		}
		fmt.Println()
	}
}

func main() {
	csll := &CircularSinglyLinkedList{}
	csll.InsertAtEnd(10)
	csll.InsertAtEnd(20)
	csll.InsertAtEnd(30)
	fmt.Println("Circular Singly Linked List after inserting at the end:")
	csll.Display()

	csll.InsertAtBegin(5)
	fmt.Println("\nCircular Singly Linked List after inserting at the beginning:")
	csll.Display()

	csll.DeleteFromBegin()
	fmt.Println("\nCircular Singly Linked List after deletion from beginning:")
	csll.Display()

	csll.DeleteFromEnd()
	fmt.Println("\nCircular Singly Linked List after deletion from end:")
	csll.Display()

	csll.DeleteTail()
	fmt.Println("\nCircular Singly Linked List after deletion from tail:")
	csll.Display()
}