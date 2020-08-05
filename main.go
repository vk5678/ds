package main

import (
	"fmt"
	dl "gofiles/datastructures/DoubleList"
	"strconv"
)

func main() {
	var doubleList dl.DoublyLinkedList = &dl.List{}
	for i := 0; i < 10; i++ {
		doubleList.Insert("m" + strconv.Itoa(i))
	}
	value, ok := doubleList.RemoveAtValue("m9")
	doubleList.RemoveAtValue("m5")

	if value == nil {
		fmt.Println(ok, dl.Head.Data, dl.Tail.Data)
	}
	if doubleList.IsEmpty() == false {
		doubleList.Display()
	}
}
