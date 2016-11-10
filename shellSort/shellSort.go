package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	length := len(s1)
	gap := int(length / 2)
	for gap >= 1 {
		for i := gap; i < length; i++ {
			value := s1[i]
			for j := i; (j-gap) >= 0 && value < s1[j-gap]; s1[j] = value {
				s1[j] = s1[j-gap]
				j = j - gap
				fmt.Printf("Sorting ...:\t%v\n", s1)
			}
		}
		gap = int(gap / 2)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
