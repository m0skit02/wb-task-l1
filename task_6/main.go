package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("1. По условию")
	stopByCondition()

	fmt.Println("2. Через канал")
	stopByChannel()

	fmt.Println("3. Через context")
	stopByContext()

	fmt.Println("4. Через runtime.Goexit")
	stopByGoExit()

	fmt.Println("5. Через таймер")
	stopByTimer()
}

func stopByCondition() {
	done := false

	go func() {
		for i := 1; ; i++ {
			if done {
				fmt.Println("Горутина завершена по условию")
				return
			}
			fmt.Printf("Горутина работает: %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	done = true
	time.Sleep(1 * time.Second)
}

func stopByChannel() {
	stopCh := make(chan struct{})

	go func() {
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				fmt.Println("Горутина завершена по каналу")
				return
			default:
				fmt.Printf("Горутина работает: %d\n", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	stopCh <- struct{}{}
	time.Sleep(1 * time.Second)
}

func stopByContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена по контексту")
				return
			default:
				fmt.Printf("Горутина работает: %d\n", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func stopByGoExit() {
	go func() {
		defer fmt.Println("Горутина завершена по runtime.Goexit")
		fmt.Println("Горутина начала работать")
		time.Sleep(1 * time.Second)
		fmt.Println("Горутина работает")
		runtime.Goexit()
		fmt.Println("Что-то")
	}()

	time.Sleep(2 * time.Second)
}

func stopByTimer() {
	go func() {
		timer := time.After(2 * time.Second)
		for i := 1; ; i++ {
			select {
			case <-timer:
				fmt.Println("Горутина завершена по таймеру")
				return
			default:
				fmt.Printf("Горутина работает: %d\n", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
}
