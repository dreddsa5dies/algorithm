package main

import "fmt"

func main() {
	s1 := []int{52, 42, 10, 32, 11, 24, 23, 11, 2423, 22, 123, 43, 87, 5, -12, 54, -1000, 1000, 1012, 32, 55, 66, 77} // срез int
	// by 2 element to sorting need gap = 0 !!
	s2 := []int{23, 12}
	shellSort(s1)
	shellSort(s2)
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
