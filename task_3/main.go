package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d получил: %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров, например: go run main.go 3")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное количество воркеров")
		return
	}

	jobs := make(chan int)

	for i := 1; i <= n; i++ {
		go worker(i, jobs)
	}

	for i := 1; ; i++ {
		jobs <- i
		time.Sleep(time.Second)
	}
}
