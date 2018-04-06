package main

import (
	"fmt"

	"os"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("List:\t%v\n", s1)
	a := util.Integer("Inter number")
	linearSearch(s1, a)
}

func linearSearch(list []int, item int) {
	for index, value := range list {
		if value == item {
			fmt.Printf("Position:\t%v\n", index)
			os.Exit(0)
		}
	}
	fmt.Println("Not found")
}
