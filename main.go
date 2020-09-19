package main

import (
	dl "ds/datastructures/DoubleList"
	st "ds/datastructures/stack"
)

func main() {
	var doublelist dl.DoublyLinkedList = &dl.List{}
	doublelist.Insert(1)
	doublelist.Insert(3)
	doublelist.Display()

	var stack st.Stack = &st.S{}
	stack.Initialize(5)
	stack.Push(1)
	stack.Push(3)
	stack.Display()
}
