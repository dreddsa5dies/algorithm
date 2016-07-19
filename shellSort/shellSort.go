package main

import "fmt"

func main() {
	s1 := []int{12, 42, 10, 32, 11, 24, 23, 11, 2423, 22, 123, 43, 87, 5, -12, 54, -1000, 1000, 1012, 32, 55, 66, 77} // срез int
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
