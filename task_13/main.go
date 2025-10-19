package main

import "fmt"

func main() {
	a := 1
	b := 2

	// Сложение и вычитание

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	// XOR

	a = 1
	b = 2

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	// a, b = b, a

	a = 1
	b = 2

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
