package main

import (
	"sync"
	"testing"
	"time"
)

const SleepDuration = time.Microsecond

type ReadAndWrite interface {
	Read()
	Write()
}

type CounterWithMutex struct {
	mu    sync.Mutex
	count int
}

func (s *CounterWithMutex) Read() {
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = s.count
	time.Sleep(SleepDuration)
}

func (s *CounterWithMutex) Write() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
	time.Sleep(SleepDuration)
}

type CounterWithRWMutex struct {
	mu    sync.RWMutex
	count int
}

func (s *CounterWithRWMutex) Read() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_ = s.count
	time.Sleep(SleepDuration)
}

func (s *CounterWithRWMutex) Write() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
	time.Sleep(SleepDuration)
}

func benchmark(b *testing.B, obj ReadAndWrite, rt, wt int) {
	for range b.N {
		var wg sync.WaitGroup

		// read op
		for range rt * 100 {
			wg.Add(1)
			go func() {
				obj.Read()
				wg.Done()
			}()
		}

		// write op
		for range wt * 100 {
			wg.Add(1)
			go func() {
				obj.Write()
				wg.Done()
			}()
		}

		wg.Wait()
	}
}

// NOTE: 读写操作次数为 9:1
func BenchmarkCounterWithMutex_RW_91(b *testing.B)   { benchmark(b, &CounterWithMutex{}, 9, 1) }
func BenchmarkCounterWithRWMutex_RW_91(b *testing.B) { benchmark(b, &CounterWithRWMutex{}, 9, 1) }

// NOTE: 读写操作次数为 1:9
func BenchmarkCounterWithMutex_RW_19(b *testing.B)   { benchmark(b, &CounterWithMutex{}, 1, 9) }
func BenchmarkCounterWithRWMutex_RW_19(b *testing.B) { benchmark(b, &CounterWithRWMutex{}, 1, 9) }

// NOTE: 读写操作次数为 5:5
func BenchmarkCounterWithMutex_RW_55(b *testing.B)   { benchmark(b, &CounterWithMutex{}, 5, 5) }
func BenchmarkCounterWithRWMutex_RW_55(b *testing.B) { benchmark(b, &CounterWithRWMutex{}, 5, 5) }
