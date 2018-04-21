package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	arr := util.RandomInt() // срез int

	fmt.Println("Unsorted List:", arr)

	k := getK(arr)
	tmpArr := make([]int, k)

	for i := 0; i < len(arr); i++ {
		tmpArr[arr[i]]++
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if tmpArr[i] > 0 {
				arr[j] = i
				j++
				tmpArr[i]--
				continue
			}
			break
		}
	}

	fmt.Println("Sorted List: ", arr)
}

func getK(arr []int) int {
	if len(arr) == 0 {
		return 1
	}

	k := arr[0]

	for _, v := range arr {
		if v > k {
			k = v
		}
	}

	return k + 1
}
