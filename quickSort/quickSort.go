package main

import "fmt"

var index, start, end int

func main() {
	s1 := []int{52, 42, 10, 32, 11, 24, 23, 11, 2423, 22, 123, 43, 87, 5, -12, 54, -1000, 1000, 1012, 32, 55, 66, 77} // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	quickSort(s1, 0, 9)
}

func partition(s1 []int, start, end int) int {
	pivot := s1[end]
	length := len(s1)
	index = start
	for current := start; current < length; current++ {
		if s1[current] <= pivot {
			s1[index], s1[current] = s1[current], s1[index]
			index++
		}
	}
	s1[index], s1[end] = s1[end], s1[index]
	fmt.Printf("Sorted list:\t%v\n", s1)
	return index
}

func quickSort(s1 []int, start, end int) {
	if start < end {
		index = partition(s1, start, end)
	}
	quickSort(s1, start, index-1)
	quickSort(s1, index+1, end)
}
