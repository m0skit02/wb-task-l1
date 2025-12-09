package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := new(big.Int)
	b := new(big.Int)

	fmt.Print("Введите число a: ")
	aStr, _ := reader.ReadString('\n')
	a.SetString(aStr, 10)

	fmt.Print("Введите число b: ")
	bStr, _ := reader.ReadString('\n')
	b.SetString(bStr, 10)

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int)

	if b.Cmp(big.NewInt(0)) != 0 {
		div.Div(a, b)
	} else {
		div = nil
	}

	fmt.Println("\nРезультаты:")
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", sub)
	fmt.Println("a * b =", mul)

	if div != nil {
		fmt.Println("a / b =", div)
	} else {
		fmt.Println("a / b = ошибка: деление на ноль")
	}
}
