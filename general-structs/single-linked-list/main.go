package main

import (
	"fmt"
)

func main() {
	list := newList()
	list.addNode(20)
	list.addNode(29)
	list.addNode(30)
	list.addNode(10)
	list.addNode(0)

	contains := list.contains(31)
	fmt.Println(contains)

	deleted := list.deleteWithValue(26)
	fmt.Println(deleted)

	list.printList()

}

type node struct {
	data any
	next *node
}

type list struct {
	size int
	head *node
}

func newList() *list {
	return &list{size: 0, head: nil}
}

func (s *list) addNode(data any) {
	if s.head == nil {
		node := node{data: data}
		s.head = &node
		s.size++
	} else {
		node := node{data: data}
		node.next = s.head
		s.head = &node
		s.size++
	}
}

func (s *list) printList() {
	if s.head == nil {
		return
	}

	currentNode := s.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (s *list) deleteWithValue(data any) bool {
	if s.size == 0 {
		return false
	}

	if s.head.data == data {
		s.head = s.head.next
		s.size--
		return true
	}

	prevNode := s.head
	for prevNode.next.data != data {
		if prevNode.next.next == nil {
			return false
		}
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
	s.size--
	return true
}

func (s *list) contains(data any) bool {
	if s.size == 0 {
		return false
	}

	if s.head.data == data {
		return true
	}

	prevNode := s.head
	for prevNode.next.data != data {
		if prevNode.next.next == nil {
			return false
		}
		prevNode = prevNode.next
	}
	return true
}
