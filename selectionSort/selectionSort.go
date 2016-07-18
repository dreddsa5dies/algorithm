package main

import "fmt"

func main() {
	s1 := []int{12, 42, 10, 32, 11, 24, 56, 23, 54, 1, -23, 200, 111, 2423, 22, 123, 43, 87, 5, -12, 54, 1000, 1012, 32, 55, 66, 77} // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
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
	fmt.Printf("Sorted list:\t%v\n", s1)
}
