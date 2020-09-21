package Queue

import (
	dl "ds/datastructures/DoubleList"
	"errors"
	"fmt"
)

//Queue ... interface implementing methods of queue
type Queue interface {
	Initialize(value int) error

	Enqueue(value interface{}) error

	Dequeue() (interface{}, error)

	Display()

	IsEmpty() bool

	IsFull() bool
}

var (
	//Size ... Size of the Queue
	Size int = 0
	//Top ... TopNode of the Queue
	TopNode *dl.List = nil
	//incrementer ... just intermidiate variable for calculating the overflow and underflow of the queue
	incrementer int = 0
	//Tail ... TailNode of the Queue
	TailNode *dl.List = nil
	//ErrQueueIsEmpty ... error describing that the queue is empty
	ErrQueueIsEmpty error = errors.New("the queue is empty")
	//ErrQueueisFull ... error describing that the queue is full
	ErrQueueIsFull error = errors.New("the queue is full")
	//ErrQueueIsAlreadyInitialized ... error describing that the queue is already initialized
	ErrQueueIsAlreadyInitialized error = errors.New("the queue has already been intialized")
)

//Q ... structure of Queue
type Q struct {
	container dl.DoublyLinkedList
}

//Initialize ... Initializing the queue
func (q *Q) Initialize(value int) error {
	if incrementer > 0 {
		return ErrQueueIsAlreadyInitialized
	}
	Size = value
	q.container = &dl.List{}
	return nil
}

//Enqueue ... inserting an element into the queue
func (q *Q) Enqueue(value interface{}) error {
	if incrementer >= Size {
		return ErrQueueIsFull
	}

	q.container.Insert(value)
	TopNode = dl.Head
	TailNode = dl.Tail
	incrementer++
	return nil
}

//Dequeue ... removing an element into the queue
func (q *Q) Dequeue() (interface{}, error) {
	if incrementer <= 0 {
		return nil, ErrQueueIsEmpty
	}
	if v, err := q.container.RemoveAtHead(); err == nil {
		TailNode = dl.Tail
		TopNode = dl.Head
		incrementer--
		return v, nil
	} else {

		return nil, err
	}
}

//Display ... displaying the elements in the queue
func (q *Q) Display() {
	if incrementer <= 0 {
		fmt.Println(ErrQueueIsEmpty)
	} else {
		fmt.Println(incrementer)
		q.container.Display()
	}
}

//IsEmpty ... returns true if the queue is empty
func (q *Q) IsEmpty() bool {
	if incrementer <= 0 {
		return true
	}
	return false
}

//IsFull ... returns true if the queue is full
func (q *Q) IsFull() bool {
	if incrementer >= Size {
		return true
	}
	return false
}
