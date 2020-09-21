package datastructures

import (
	"errors"
	"fmt"
)

var (
	//Head ... Head of the List
	Head *List = nil
	//Tail ... Tail of the List
	Tail *List = nil
	//Size ... Size of the List
	Size int = 0
	//ErrListIsEmpty ... error message for an empty List
	ErrListIsEmpty error = errors.New("the List is Empty")
	//ErrElementNotFound ... error message for element not found
	ErrElementNotFound = errors.New("the Node with the given key is not found")
)

//DoublyLinkedList ... interface implementing structure List and its methods
type DoublyLinkedList interface {
	Insert(value interface{})
	InsertAfterValue(value interface{}, key interface{})
	InsertAtHead(value interface{})
	Remove() (interface{}, error)
	RemoveAtValue(value interface{}) (interface{}, error)
	RemoveAtHead() (interface{}, error)
	Display() error
	IsEmpty() bool
}

//List ... Structure of the List
type List struct {
	Data interface{}
	Next *List
	Prev *List
}

//IsEmpty ... returns true if List is empty;else returns false
func (l *List) IsEmpty() bool {
	if Size == 0 || Head == nil && Tail == nil {
		return true
	}
	return false
}

//Insert ... inserts elements of any type to the Tail of the List
func (l *List) Insert(value interface{}) {
	if Head == nil {
		newNode := new(List)
		newNode.Data = value
		newNode.Next = nil
		newNode.Prev = nil
		Head = newNode
		Tail = newNode
		Size++
	} else {
		Tail.Next = &List{value, nil, Tail}
		Tail = Tail.Next
		Size++

	}
}

//InsertAtHead ... inserts elements at the head of the List
func (l *List) InsertAtHead(value interface{}) {
	if Head == nil {
		newNode := &List{value, nil, nil}
		Head = newNode
		Tail = newNode
		Size++
	} else {
		newNode := &List{value, nil, nil}
		newNode.Next = Head
		newNode.Prev = nil
		Head.Prev = newNode
		Head = newNode
		Size++
	}
}

//InsertAfterValue ... inserts elements after the given key value
func (l *List) InsertAfterValue(value interface{}, key interface{}) {
	if Head == nil {
		newNode := &List{value, nil, nil}
		Head = newNode
		Tail = newNode
		Size++
	} else if Head.Data == key {
		newNode := &List{value, nil, nil}
		Head.Next = newNode
		newNode.Prev = Head
		Tail = newNode
		Size++
	} else {
		current := Head
		for current != nil {
			if current.Data == key {

				newNode := &List{value, nil, nil}
				current.Next = newNode
				newNode.Prev = current
				newNode.Next = current.Next.Next
				if newNode.Next == nil {
					Tail = newNode
				}
				Size++
				break
			}
			current = current.Next
		}
	}
}

//Remove ... removes the elements at the tail of the list
func (l *List) Remove() (interface{}, error) {
	if Tail == nil {
		return nil, ErrListIsEmpty
	}
	if Size == 1 {
		current := Head
		Head = nil
		Tail = nil
		Size--
		return current, nil

	}
	before := Tail.Prev
	before.Next = nil
	Tail.Prev = nil
	returner := Tail
	Tail = before
	Size--
	return returner, nil
}

//RemoveAtHead ... removes the elements at the Head
func (l *List) RemoveAtHead() (interface{}, error) {
	if Head == nil {
		return nil, ErrListIsEmpty
	}
	if Size == 1 {
		temp := Head
		Head = nil
		Tail = nil
		Size--
		return temp, nil

	}
	temp := Head
	Head = Head.Next
	Head.Prev = nil
	Size--
	return temp, nil

}

//RemoveAtValue ... removes the element at certain value
func (l *List) RemoveAtValue(value interface{}) (interface{}, error) {
	if l.IsEmpty() == true {
		return nil, ErrListIsEmpty
	}
	current := Head
	for current != nil {
		if current.Data == value {

			if current.Next != nil {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev

				Size--
				return current, nil
			}
			current.Prev.Next = nil
			Tail = current.Prev
			current.Prev = nil

			Size--
			return current, nil

		}
		current = current.Next
	}
	return nil, ErrElementNotFound
}

//Display ... displays the elements of the List if it's not empty
func (l *List) Display() error {
	if Head == nil && Size == 0 {
		return ErrListIsEmpty
	}

	current := Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
	return nil
}
