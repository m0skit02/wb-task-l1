package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	// Убираем возможный перевод строки
	input = input[:len(input)-1]

	result := reverseString(input)
	fmt.Println("Перевёрнутая строка:", result)
}
