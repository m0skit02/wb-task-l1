package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	intersection := intersect(A, B)
	fmt.Println("Пересечение =", intersection)
}

func intersect(a, b []int) []int {
	set := make(map[int]bool)
	var result []int

	for _, v := range a {
		set[v] = true
	}

	for _, v := range b {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}

	return result
}
