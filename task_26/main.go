package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	seen := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}

func main() {
	fmt.Println(hasUniqueChars("abcd"))
	fmt.Println(hasUniqueChars("abCdefAaf"))
	fmt.Println(hasUniqueChars("aabcd"))
}
