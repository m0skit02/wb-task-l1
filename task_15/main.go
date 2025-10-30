package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	v := createHugeString(1 << 10)
	return string([]byte(v[:100]))
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}
