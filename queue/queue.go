package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
	"github.com/dreddsa5dies/algorithm/util/queue"
)

func main() {
	list := util.RandomInt() // срез int
	fmt.Printf("List:\t%v\n", list)

	q := queue.QueueNew()
	fmt.Println("Len Queue: ", q.Len())

	fmt.Println("Enqueue:")
	for i := 0; i < len(list); i++ {
		fmt.Print(list[i], " ")
		q.Enqueue(list[i])
	}
	fmt.Println("")

	fmt.Println("Len Queue: ", q.Len())
	fmt.Println("Peek Queue: ", q.Peek())

	fmt.Println("Dequeue:")
	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}
