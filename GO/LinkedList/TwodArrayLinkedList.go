package main

import "fmt"


type Node struct {
	data [][]int
	row int
	col int
	next *Node
}

type TwodArrayLinkedList struct {
	head    *Node
	tail    *Node
}

func (all *TwodArrayLinkedList) InsertAtEnd(data [][]int, row, col int) {
	newNode := &Node{data: data, row: row, col: col}
	if all.head == nil {
		all.head = newNode
		all.tail = newNode
	} else {
		all.tail.next = newNode
		all.tail = newNode
	}
}

func (lll *TwodArrayLinkedList) display(){
	if lll.head == nil {
		fmt.Println("List is empty.")
	} else {
		current := lll.head
		for current != nil {
			// fmt.Println(current.data, "Rows: ", current.row, " Cols: ", current.col)
			for i := 0; i < current.row; i++ {
				for j := 0; j < current.col; j++ {
					fmt.Print(current.data[i][j], " ")
				}
				fmt.Println()
			}
			if (current.next != nil){
				fmt.Println("|")
				fmt.Println("v")
			}
			current = current.next
		}
	}
}


func main() {
	all := &TwodArrayLinkedList{}
	var TwodArrayLinkedListSize int
	fmt.Print("Enter the size of Linked List: ")
	fmt.Scan(&TwodArrayLinkedListSize)

	for i := 0; i < TwodArrayLinkedListSize; i++ {
		var row, col int
		fmt.Printf("\nEnter the number of rows for node %d: ", i+1)
		fmt.Scan(&row)
		fmt.Printf("Enter the number of columns for node %d: ", i+1)
		fmt.Scan(&col)

		data := make([][]int, row)
		for j := 0; j < row; j++ {
			data[j] = make([]int, col)
		}
		for j := 0; j < row; j++ {
			for k := 0; k < col; k++ {
				fmt.Printf("Enter element for row %d, column %d: ", j+1, k+1)
				fmt.Scan(&data[j][k])
			}
		}
		all.InsertAtEnd(data, row, col)
	}
	fmt.Println()
	all.display()
}