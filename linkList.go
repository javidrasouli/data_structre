package main

import (
	"fmt"
)

type Node struct {
	data   any
	Next   *Node
	Before *Node
}

type LinkList struct {
	Head *Node
	tail *Node
}

func NewList() *LinkList {
	return new(LinkList)
}

func NewNode(data any) *Node {
	return &Node{
		data: data,
	}
}

func (l *LinkList) Insert(data any) {
	node := NewNode(data)
	if l.Head == nil {
		l.Head = node
		l.tail = node
		return
	}
	beforeNode := l.tail
	l.tail.Next = node
	l.tail = node
	l.tail.Before = beforeNode
}

func (l *LinkList) Length() int {
	counter := 0
	node := l.Head
	for node.Next == nil {
		node = node.Next
		counter++
	}
	return counter
}

func (l *LinkList) UpdateAtIndex(index int, data any) {
	counter := 0
	node := l.Head
	if index < 0 {
		fmt.Println("index out of rang")
		return
	}
	if index == 0 {
		l.Head.data = data
	}
	for node.Next == nil {
		if counter == index {
			node.data = data
		}
	}
	fmt.Println("data changed")
}

func (l *LinkList) Print() {
	node := l.Head
	for node.Next == nil {
		fmt.Println(node.data)
		node = node.Next
	}
}
func (l *LinkList) DeleteAtIndex(index int) {
	node := l.Head
	i := 0
	if index < 0 {
		fmt.Println("index out of range")
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		return
	}
	for node != nil {
		node = node.Next
		if i == index-1 {
			node = node.Next.Next
		}
		i++
	}
	fmt.Println("index deleted")
}

func (l *LinkList) Perpend(data any) {
	node := NewNode(data)
	l.Head.Before = node
	node.Next = l.Head
	l.Head = node
	return
}

func (l *LinkList) DeleteFirst() {
	l.Head.Next.Before = nil
	l.Head = l.Head.Next
}

func (l *LinkList) DeleteLast() {
	l.tail.Before.Next = nil
}

func (l *LinkList) InsertAtIndex(index int, data any) {
	counter := 0
	node := l.Head
	nNode := NewNode(data)
	for node.Next == nil {
		if counter == index-1 {
			node.Next = nNode
			nNode.Before = node
		}
		if counter == index {
			nNode.Next = node.Next.Next
			node = nNode
		}

		node = node.Next

		counter++
	}
}

func (l *LinkList) FindAtIndex(index int) {
	counter := 0

	node := l.Head

	for node.Next == nil {

		if counter == index {
			fmt.Println(node.data)
			return
		}
		counter++
		node = node.Next
	}
}

func (l *LinkList) FindFirstIndex(val any) int {
	counter := -1
	node := l.Head

	for node.Next == nil {
		if node.data == val {
			return counter + 1
		}
		node = node.Next
		counter++
	}
	return counter
}

func (l *LinkList) FindLastIndex(val any) int {
	counter := -1
	lastFind := -1
	node := l.Head

	for node.Next == nil {
		if node.data == val {
			lastFind = counter + 1
		}
		counter++
		node = node.Next
	}

	return lastFind
}

func main() {
	list := NewList()
	list.Insert("a")
	list.Insert("b")
	list.Insert("c")
	list.Insert("d")

}
