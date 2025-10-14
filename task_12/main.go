package main

import "fmt"

func main() {
	// Исходная последовательность
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём множество (map без повторений)
	set := make(map[string]struct{})

	for _, word := range words {
		set[word] = struct{}{}
	}

	// Выводим множество
	fmt.Println("Множество уникальных слов:")
	for word := range set {
		fmt.Println(word)
	}
}
