package main

import (
	"fmt"
)

func reverseRunes(r []rune, l, rgt int) {
	for l < rgt {
		r[l], r[rgt] = r[rgt], r[l]
		l++
		rgt--
	}
}

func reverseWords(s string) string {
	r := []rune(s)

	reverseRunes(r, 0, len(r)-1)

	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverseRunes(r, start, i-1)
			start = i + 1
		}
	}

	return string(r)
}

func main() {
	input := "snow dog sun"
	result := reverseWords(input)
	fmt.Println(result)
}
