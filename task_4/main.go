package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт, выходим\n", id)
				return
			}
			fmt.Printf("Worker %d обработал: %d\n", id, job)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров, например: go run . 3")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное количество воркеров")
		return
	}

	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- i:
				time.Sleep(time.Second)
			}
		}
	}()

	<-sigCh
	fmt.Println("\nПолучен сигнал SIGINT, завершаем работу...")
	cancel()

	wg.Wait()
	fmt.Println("Все воркеры завершились. Программа остановлена.")
}
