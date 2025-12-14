package main

import (
	"fmt"
)

func removeAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}

	copy(s[i:], s[i+1:])

	var zero T
	s[len(s)-1] = zero

	return s[:len(s)-1]
}

func main() {
	type Data struct {
		Value string
	}

	d1 := &Data{Value: "one"}
	d2 := &Data{Value: "two"}
	d3 := &Data{Value: "three"}
	d4 := &Data{Value: "four"}

	s := []*Data{d1, d2, d3, d4}

	fmt.Println("До удаления:")
	for i, v := range s {
		fmt.Printf("%d: %s\n", i, v.Value)
	}

	s = removeAt(s, 1)

	fmt.Println("\nПосле удаления:")
	for i, v := range s {
		fmt.Printf("%d: %s\n", i, v.Value)
	}
}
