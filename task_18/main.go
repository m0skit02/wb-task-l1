package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterMutex struct {
	mu sync.Mutex
	n  int
}

func (c *CounterMutex) IncMutex() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *CounterMutex) ValueMutex() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

type CounterAtomic struct {
	n int64
}

func (c *CounterAtomic) IncAtomic() {
	atomic.AddInt64(&c.n, 1)
}

func (c *CounterAtomic) ValueAtomic() int64 {
	return atomic.LoadInt64(&c.n)
}

func main() {
	var wg sync.WaitGroup
	counterMutex := &CounterMutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterMutex.IncMutex()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счётчика: (Mutex)", counterMutex.ValueMutex())

	counterAtomic := &CounterAtomic{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterAtomic.IncAtomic()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счётчика: (Atomic)", counterAtomic.ValueAtomic())
}
