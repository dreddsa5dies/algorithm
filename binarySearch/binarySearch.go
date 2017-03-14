package main

import (
	"fmt"

	"sort"

	"os"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	sort.Ints(s1)
	fmt.Printf("Sorted list:\t%v\n", s1)
	a := util.Integer("Inter number")
	binSearch(s1, a)
}

func binSearch(list []int, item int) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high)
		guess := list[mid]
		if guess == item {
			fmt.Printf("Position:\t%v\n", mid)
			os.Exit(1)
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Not found")
}
