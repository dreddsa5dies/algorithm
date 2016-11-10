package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

var index, start, end int

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	sort(s1, 0, len(s1)-1)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}

func sort(s1 []int, start, end int) {
	if start >= end {
		return
	}

	pivot := s1[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > s1[j] {
			s1[i], s1[j] = s1[j], s1[i]
			i++
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}

	s1[start], s1[i-1] = s1[i-1], s1[start]

	sort(s1, start, i-2)
	sort(s1, i, end)
}
