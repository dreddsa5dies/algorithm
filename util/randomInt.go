package util

import "math/rand"

// RandomInt create random array []int
func RandomInt() []int {
	list := rand.Perm(20)
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}
