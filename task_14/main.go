package main

import (
	"fmt"
)

// Функция для определения типа переменной
func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int:
		fmt.Println("Тип: chan int")
	case chan string:
		fmt.Println("Тип: chan string")
	case chan bool:
		fmt.Println("Тип: chan bool")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)
	detectType("привет")
	detectType(true)

	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	detectType(ch1)
	detectType(ch2)
	detectType(ch3)

	var unknown float64 = 3.14
	detectType(unknown)
}
