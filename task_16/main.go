package main

import (
	"fmt"
)

// quickSort выполняет быструю сортировку массива целых чисел
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // базовый случай: массив длиной 0 или 1 уже отсортирован
	}

	// Выбираем опорный элемент (pivot) — можно взять середину
	pivot := arr[len(arr)/2]

	// Разделяем элементы на три группы
	var left, equal, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			right = append(right, v)
		}
	}

	// Рекурсивно сортируем левую и правую часть
	left = quickSort(left)
	right = quickSort(right)

	// Объединяем части
	return append(append(left, equal...), right...)
}

func main() {
	arr := []int{9, 3, 7, 1, 8, 2, 5, 6, 4}
	fmt.Println("До сортировки:", arr)

	sorted := quickSort(arr)
	fmt.Println("После сортировки:", sorted)
}
