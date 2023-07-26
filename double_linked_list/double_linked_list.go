package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func InitDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) AddFront(val int) {
	newNode := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}

	l.Len++
}

func (l *DoubleLinkedList) AddBack(val int) {
	newNode := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		newNode.Prev = current
		current.Next = newNode
		l.Tail = newNode
	}

	l.Len++
}

func (l *DoubleLinkedList) RemoveFront() error {
	if l.Head == nil {
		return fmt.Errorf("cannot remove empty linked list")
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	}

	l.Len--
	return nil
}

func (l *DoubleLinkedList) RemoveBack() error {
	if l.Tail == nil {
		return fmt.Errorf("cannot remove empty linked list")
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	}

	l.Len--
	return nil
}

func (l *DoubleLinkedList) TraverseForward() error {
	if l.Head == nil {
		return fmt.Errorf("empty linked list")
	}

	temp := l.Head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Val, temp.Prev, temp.Next)
		temp = temp.Next
	}
	fmt.Println()

	return nil
}

func (l *DoubleLinkedList) TraverseReverse() error {
	if l.Head == nil {
		return fmt.Errorf("empty linked list")
	}

	temp := l.Tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Val, temp.Prev, temp.Next)
		temp = temp.Prev
	}
	fmt.Println()

	return nil
}

func (l *DoubleLinkedList) PrintAll() error {
	if l.Head == nil {
		return fmt.Errorf("empty linked list")
	}

	temp := l.Head
	for temp != nil {
		fmt.Printf("%d | ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
	return nil
}

func (l *DoubleLinkedList) GetLength() int {
	return l.Len
}

func main() {
	dll := InitDoubleLinkedList()
	dll.AddFront(2)
	dll.AddFront(4)
	dll.AddFront(6)
	dll.AddFront(8)
	dll.AddFront(10)
	dll.AddFront(12)
	dll.RemoveFront()
	dll.RemoveBack()
	dll.PrintAll()
	dll.TraverseForward()
}
