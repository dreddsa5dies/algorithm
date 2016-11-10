package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	shellSort(s1)
}

func shellSort(s1 []int) {
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	length := len(s1)
	gap := int(length / 3) // modify gap by 3
	if gap == 0 {
		gap = 1
	}
	for gap >= 1 {
		for i := gap; i < length; i++ {
			value := s1[i]
			for j := i; (j-gap) >= 0 && value < s1[j-gap]; s1[j] = value {
				s1[j] = s1[j-gap]
				j = j - gap
			}
		}
		fmt.Printf("gap:\t%v\n", gap)
		gap = int(gap / 3) // modify gap by 3
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
