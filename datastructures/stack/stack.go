package stack

import (
	"errors"
	"fmt"
)

type Stack interface {
	Initialize(value int) error
	Push(value interface{}) (bool, error)
	Pop() (interface{}, error)
	Display()
}

var (
	Top             int   = -1
	Size            int   = -1
	ErrStackIsFull  error = errors.New("The stack is Full.")
	ErrStackIsEmpty error = errors.New("The stack is Empty")
)

type S struct {
	container []interface{}
}

func (s *S) Initialize(value int) error {
	if Top == -1 {
		Size = value
		s.container = make([]interface{}, 0)
		return nil
	} else {
		return ErrStackIsFull
	}
}
func (s *S) Push(value interface{}) (bool, error) {
	if Top >= Size {
		return false, ErrStackIsFull
	} else {
		s.container = append(s.container, value)
		Top++
		return true, nil
	}
}
func (s *S) Pop() (interface{}, error) {
	if Top <= -1 {
		return nil, ErrStackIsEmpty
	} else {
		v := Top
		val := s.container[v]
		s.container[v] = nil
		Top--
		return val, nil
	}
}
func (s *S) Display() {
	if Top == -1 {
		fmt.Println(ErrStackIsEmpty)
	} else {
		for index, value := range s.container {
			fmt.Printf("%d=>%d\n", index, value)
		}
		fmt.Printf("The Size of the stack is %d\n", Size)
		fmt.Printf("The top of the stack is %d\n", s.container[Top])
	}
}
