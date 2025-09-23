package main

import (
	"fmt"
	"strconv"
)

func setBit(x int64, i uint, val int) int64 {
	if val == 1 {
		x |= 1 << i
	} else {
		x &^= 1 << i
	}
	return x
}

func formatBinary(x int64) string {
	bin := strconv.FormatInt(x, 2)
	width := len(bin)
	if width < 4 {
		width = 4
	}
	return fmt.Sprintf("%0*b", width, x)
}

func main() {
	var x int64
	var i uint
	var val int

	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	fmt.Print("Введите номер бита (начиная с 0): ")
	fmt.Scan(&i)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&val)

	fmt.Printf("Исходное число: %d -> %s\n", x, formatBinary(x))

	x = setBit(x, i, val)

	fmt.Printf("Результат: %d -> %s\n", x, formatBinary(x))
}
