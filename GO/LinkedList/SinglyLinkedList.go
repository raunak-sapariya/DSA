package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (sll *SinglyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if sll.head == nil {
		sll.head = newNode
	} else {
		current := sll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (sll *SinglyLinkedList) InsertAtBegin(data int) {
	newNode := &Node{data: data}
	if sll.head == nil {
		sll.head = newNode
	} else {
		newNode.next = sll.head
		sll.head = newNode
	}
}

func (sll *SinglyLinkedList) DeleteFromEnd() {
	if sll.head == nil {
		return
	} else {
		if sll.head.next == nil {
			sll.head = nil
		} else {
			current := sll.head
			for current.next.next != nil {
				current = current.next
			}
			current.next = nil
		}
	}	
}

func (sll *SinglyLinkedList) DeleteFromBegin() {
	if sll.head == nil {
		return
	} else {
		sll.head = sll.head.next
	}
}

func (sll *SinglyLinkedList) Display() {
	if sll.head == nil {
		fmt.Println("List is empty.")
	} else {
		current := sll.head
		for current != nil {
			fmt.Print(current.data, " -> ")
			current = current.next
		}
		fmt.Println("nil")
	}
}

func main() {
	sll := &SinglyLinkedList{}
	sll.InsertAtEnd(10)
	sll.InsertAtEnd(20)
	sll.InsertAtEnd(30)
	fmt.Println("Singly Linked List after inserting at the end:")
	sll.Display()

	sll.InsertAtBegin(5)
	fmt.Println("\nSingly Linked List after inserting at the beginning:")
	sll.Display()

	sll.DeleteFromEnd()
	fmt.Println("\nSingly Linked List after deletion from end:")
	sll.Display()
  
	sll.DeleteFromBegin()
	fmt.Println("\nSingly Linked List after deletion from beginning:")
	sll.Display()
}