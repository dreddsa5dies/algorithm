package main

import (
	"dreddsa5dies/algorithm/util"
	"fmt"
)

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	for i := 1; i < len(s1); i++ {
		for j := i; j != 0 && s1[j] < s1[j-1]; j-- {
			s1[j-1], s1[j] = s1[j], s1[j-1]
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
