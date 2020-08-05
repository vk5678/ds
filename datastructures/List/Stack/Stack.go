package datastructures

var (
	Size            int    = 0
	sizeIncrementer int    = 0
	Top             *Stack = nil
)

type Stack interface {
	Initialize(size int)
	push(value interface{}) error
	pop() (interface{}, error)
	IsEmpty() bool
	IsFull() bool
}
type element struct {
	Value interface{}
	Next  *Stack
	Prev  *Stack
}

func (e *Element) Initialize(size)
