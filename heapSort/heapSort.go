package main

import "fmt"

var maxChild int

func main() {
	s1 := []int{12, 42, 10, 32, 11, 24, 23, 11, 2423, 22, 123, 43, 87, 5, -12, 54, -1000, 1000, 1012, 32, 55, 66, 77} // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")

	i := 0
	tmp := 0

	for i = len(s1)/2 - 1; i >= 0; i-- {
		s1 = heapSort(s1, i, len(s1))
	}

	for i = len(s1) - 1; i >= 1; i-- {
		tmp = s1[0]
		s1[0] = s1[i]
		s1[i] = tmp
		s1 = heapSort(s1, 0, i)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}

func heapSort(s1 []int, i int, s1Len int) []int {
	done := false

	tmp := 0
	maxChild := 0

	for (i*2+1 < s1Len) && (!done) {
		if i*2+1 == s1Len-1 {
			maxChild = i*2 + 1
		} else if s1[i*2+1] > s1[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if s1[i] < s1[maxChild] {
			tmp = s1[i]
			s1[i] = s1[maxChild]
			s1[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}

	return s1
}
