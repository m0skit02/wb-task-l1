package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SafeMap

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	v, ok := s.m[key]
	s.mu.Unlock()
	return v, ok
}

func (s *SafeMap) Len() int {
	s.mu.Lock()
	l := len(s.m)
	s.mu.Unlock()
	return l
}

func (s *SafeMap) SumValues() int {
	s.mu.Lock()
	sum := 0
	for _, v := range s.m {
		sum += v
	}
	s.mu.Unlock()
	return sum
}

// sync.Map

type SyncMapWrapper struct {
	m sync.Map
}

func (w *SyncMapWrapper) Store(key string, value int) {
	w.m.Store(key, value)
}

func (w *SyncMapWrapper) Load(key string) (int, bool) {
	v, ok := w.m.Load(key)
	if !ok {
		return 0, false
	}
	return v.(int), true
}

func (w *SyncMapWrapper) RangeSum() int {
	sum := 0
	w.m.Range(func(k, v any) bool {
		sum += v.(int)
		return true
	})
	return sum
}

func (w *SyncMapWrapper) Count() int {
	cnt := 0
	w.m.Range(func(k, v any) bool {
		cnt++
		return true
	})
	return cnt
}

func runWithSafeMap() {
	const G = 100  // горутины
	const N = 1000 // записи
	sm := NewSafeMap()
	var wg sync.WaitGroup
	wg.Add(G)

	for i := 0; i < G; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
			for j := 0; j < N; j++ {
				key := fmt.Sprintf("g%d_k%d", id, r.Intn(1000)) // немного коллизий ключей
				sm.Set(key, id+j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("SafeMap: len =", sm.Len(), "sum =", sm.SumValues())
}

func runWithSyncMap() {
	const G = 100
	const N = 1000
	var wg sync.WaitGroup
	wg.Add(G)
	w := &SyncMapWrapper{}

	for i := 0; i < G; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
			for j := 0; j < N; j++ {
				key := fmt.Sprintf("g%d_k%d", id, r.Intn(1000))
				w.Store(key, id+j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("sync.Map: len =", w.Count(), "sum =", w.RangeSum())
}

func main() {
	fmt.Println("Запуск конкурентной записи в map")
	start := time.Now()
	runWithSafeMap()
	fmt.Println("----")
	runWithSyncMap()
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
}
