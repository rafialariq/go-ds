package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

// initialization function
func InitLinkedList() *LinkedList {
	return &LinkedList{}
}

// add new node in front of linked list (new head)
func (l *LinkedList) AddFront(val int) {
	node := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}

	l.Len++

}

// add new node in the back of linked list
func (l *LinkedList) AddBack(val int) {
	node := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}

	l.Len++

}

// remove one node in front of the linked list
func (l *LinkedList) RemoveFront() error {
	if l.Head == nil {
		return fmt.Errorf("cannot remove empty linked list")
	}

	l.Head = l.Head.Next
	l.Len--
	return nil
}

// remove one node in the back of linked list
func (l *LinkedList) RemoveBack() error {
	if l.Head == nil {
		return fmt.Errorf("cannot remove empty linked list")
	}

	var prev *Node
	current := l.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}

	if prev != nil {
		prev.Next = nil
	} else {
		l.Head = nil
	}

	l.Len--
	return nil
}

// return one node value in front of the linked list
func (l *LinkedList) GetFront() (int, error) {
	if l.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}

	return l.Head.Val, nil
}

// return length of the linked list
func (l *LinkedList) GetLength() int {
	return l.Len
}

// traverse and print all of the node
func (l *LinkedList) PrintAll() error {
	if l.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%d | ", current.Val)
		current = current.Next
	}
	fmt.Println()

	return nil
}

func main() {
	ll := InitLinkedList()
	ll.AddFront(2)
	ll.AddFront(4)
	ll.AddFront(6)
	ll.AddFront(8)
	ll.PrintAll()

	ll.AddBack(4)
	ll.AddBack(6)
	ll.AddBack(8)
	ll.PrintAll()

	ll.RemoveFront()
	ll.RemoveBack()
	ll.PrintAll()
	val, _ := ll.GetFront()
	fmt.Println("First Node Value : ", val)
	len := ll.GetLength()
	fmt.Println("Linked List Length : ", len)
}
