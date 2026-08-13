package main

import "fmt"


type Node struct {
	coeff int
	pow int
	next *Node
}

type LinerLinkedList struct {
	head    *Node
	tail    *Node
}

func (lll *LinerLinkedList) create(coeff int, pow int) *Node {
	newNode := &Node{coeff: coeff, pow: pow}
	return newNode
}

func (lll *LinerLinkedList) InsertAtEnd(coeff int, pow int){
	newNode := lll.create(coeff, pow)
	if(lll.head == nil){
		lll.head = newNode
		lll.tail = newNode
	} else {
		lll.tail.next = newNode
		lll.tail = newNode
	}
}

func (lll *LinerLinkedList) display(){
	if lll.head == nil {
		fmt.Println("List is empty.")
	} else {
		current := lll.head
		for current != nil {
			fmt.Print(current.coeff, "X","^",current.pow)
			if (current.next != nil){
				fmt.Print("+")
			}
			current = current.next
		}
	}
}


func main(){
	lll := &LinerLinkedList{}
	lll.InsertAtEnd(2,5)
	lll.InsertAtEnd(5,2)
	lll.InsertAtEnd(76,2)
	lll.InsertAtEnd(8,3)

	lll.display()

}