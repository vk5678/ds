package datastructures

import (
	"fmt"
)

var (
	//Size ... Size of the list
	Size int = 0
	//Head ... Head of the List
	Head *LinkedList = nil
	//Tail ... tail of the List
	Tail *LinkedList = nil
)

//SingleLinkedList ... interface implementing the methods
type SingleLinkedList interface {
	Insert(d int)
	InsertAtHead(d int)
	InsertAfterValue(d int, value int)
	Remove() *LinkedList
	RemoveValue(value int) *LinkedList
	RemoveHead() (*LinkedList, *LinkedList)

	Display()
}

//LinkedList ... Structure of LinkedList
type LinkedList struct {
	Data int
	Next *LinkedList
}

//Insert ... method for inserting a node at end
func (l *LinkedList) Insert(d int) {
	if Size < 1 {
		Head = &LinkedList{Data: d, Next: nil}
		Tail = Head
		Size++
		return
	}
	pointer := Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	newNode := new(LinkedList)
	newNode.Data = d
	newNode.Next = nil
	pointer.Next = newNode
	Tail = newNode
	Size++
}

//InsertAtHead ... Insertion of node at Head
func (l *LinkedList) InsertAtHead(d int) {
	newNode := &LinkedList{d, nil}
	newNode.Next = Head

	Head = newNode
	Size++

}

//InsertAfterValue ... Insertion of node after a value
func (l *LinkedList) InsertAfterValue(d int, value int) {
	current := Head
	for current != nil {
		if current.Data == d {
			newNode := &LinkedList{value, nil}
			temp := &LinkedList{}
			temp = current.Next
			current.Next = newNode
			newNode.Next = temp
		}
		current = current.Next
	}

	Size++
}

//Remove ... Removal of tail Node
func (l *LinkedList) Remove() *LinkedList {
	current := Head
	tail := Head
	for current.Next != nil {
		tail = current
		current = current.Next

	}
	Tail = tail
	tail.Next = nil
	Size--
	return current

}

//RemoveValue ... Removal of node of a value
func (l *LinkedList) RemoveValue(value int) *LinkedList {
	current := Head
	before := Head
	for current != nil {

		if current.Data == value {
			before.Next = current.Next
		}
		before = current
		current = current.Next
	}
	Size--
	return current
}

//RemoveHead ... Removal of Head node
func (l *LinkedList) RemoveHead() (RemovedHead, newHead *LinkedList) {
	RemovedHead = Head
	Head = Head.Next
	newHead = Head
	Size--
	return
}

//Display ... displaying the list
func (l *LinkedList) Display() {

	current := Head
	if current != nil {
		for current != nil {
			if current == Head {
				fmt.Printf("%d->", current.Data)
				current = current.Next
			} else {
				fmt.Printf("->%d", current.Data)
				current = current.Next
			}
		}

	}
}
