package main

import (
	"fmt"
	"os"

	"sort"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	sort.Ints(s1)
	fmt.Printf("Sorted list:\t%v\n", s1)
	a := util.Integer("Inter number")
	interSearch(s1, a)
}

func interSearch(sortedArray []int, toFind int) {
	var mid int
	low := 0
	high := len(sortedArray) - 1

	for sortedArray[low] < toFind && sortedArray[high] > toFind {
		mid = low + ((toFind-sortedArray[low])*(high-low))/(sortedArray[high]-sortedArray[low])

		if sortedArray[mid] < toFind {
			low = mid + 1
		} else if sortedArray[mid] > toFind {
			high = mid - 1
		} else {
			fmt.Println("Found: ", mid)
			os.Exit(0)
		}
	}

	if sortedArray[low] == toFind {
		fmt.Println("Found: ", low)
		os.Exit(0)
	} else if sortedArray[high] == toFind {
		fmt.Println("Found: ", high)
		os.Exit(0)
	} else {
		fmt.Println("Not found")
	}
}
