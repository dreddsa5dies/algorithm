package main

import (
	"dreddsa5dies/algorithm/util"
	"fmt"
)

func main() {
	// sort on the right
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	length := len(s1)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if s1[j] > s1[j+1] {
				s1[j], s1[j+1] = s1[j+1], s1[j]
			}
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
