package main

import "fmt"

func main() {
	s1 := []int{12, 42, 10, 32, 11, 24, 23, 11, 2423, 22, 123, 43, 87, 5, -12, 54, -1000, 1000, 1012, 32, 55, 66, 77} // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	for i := 1; i < len(s1); i++ {
		value := s1[i]
		for j := i; j != 0 && value < s1[j-1]; s1[j] = value {
			s1[j] = s1[j-1]
			j--
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Printf("Sorted list:\t%v\n", s1)
}
