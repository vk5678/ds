package main

import (
	qu "ds/datastructures/Queue"
	"fmt"
)

func main() {
	var qq qu.Queue = &qu.Q{}
	qq.Initialize(6)
	qq.Enqueue(1)
	qq.Enqueue(2)
	qq.Enqueue(3)
	qq.Enqueue(5)

	fmt.Println(qu.TopNode)
	fmt.Println(qu.TailNode)
	qq.Display()

}
