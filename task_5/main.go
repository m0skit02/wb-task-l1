package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	fmt.Println("Введите время работы программы (сек):")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	ch := make(chan int)
	timer := time.After(time.Duration(N) * time.Second)

	go func() {
		i := 1
		for {
			select {
			case ch <- i:
				i++
				time.Sleep(500 * time.Millisecond)
			case <-timer:
				close(ch)
				return
			}
		}
	}()

	for val := range ch {
		fmt.Println("Принято:", val)
	}

	fmt.Println("Программа завершена после", N, "секунд")
}
