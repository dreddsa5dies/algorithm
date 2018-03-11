package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
	"github.com/dreddsa5dies/algorithm/util/stack"
)

func main() {
	list := util.RandomInt() // срез int
	fmt.Printf("List:\t%v\n", list)

	s := stack.StackNew()
	fmt.Println("Len Stack: ", s.Len())

	fmt.Println("Push:")
	for i := 0; i < len(list); i++ {
		fmt.Print(list[i], " ")
		s.Push(list[i])
	}
	fmt.Println("")

	fmt.Println("Len Stack: ", s.Len())
	fmt.Println("Peek Stack: ", s.Peek())

	fmt.Println("Pop:")
	for s.Len() != 0 {
		val := s.Pop()
		fmt.Print(val, " ")
	}
}
