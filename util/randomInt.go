package util

import (
	"fmt"
	"math/rand"
)

// RandomInt create random array []int, len()=20
func RandomInt() []int {
	list := rand.Perm(20)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// Integer ввод целого числа в stdin
func Integer(msg string) int {
	fmt.Print(msg + " > ")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return num
}
