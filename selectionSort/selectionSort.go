package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	for i := 0; i < len(s1); i++ {
		minIndex := i
		j := i + 1
		for j < len(s1) {
			if s1[j] < s1[minIndex] {
				minIndex = j
			}
			j = j + 1
		}
		s1[i], s1[minIndex] = s1[minIndex], s1[i]
		fmt.Printf("Sorting list:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
