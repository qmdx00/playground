package main

import (
	"fmt"
	"sync"
	"testing"
)

var _ sync.Locker = (*XMutex)(nil)

type XMutex struct {
	mu chan struct{}
}

func NewXMutex() *XMutex {
	return &XMutex{mu: make(chan struct{}, 1)}
}

func (x *XMutex) Lock() {
	x.mu <- struct{}{}
}

func (x *XMutex) Unlock() {
	<-x.mu
}

// NOTE: 模拟互斥锁实现
func testChannelMutexAdd(_ *testing.T, nums uint) {
	var count = 0
	var wg sync.WaitGroup
	mutex := NewXMutex()

	for range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range 10000 {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Want %d, Got %d\n", 10000*nums, count)
}

func TestChannelMutexAdd_10(t *testing.T)   { testChannelMutexAdd(t, 10) }
func TestChannelMutexAdd_100(t *testing.T)  { testChannelMutexAdd(t, 100) }
func TestChannelMutexAdd_1000(t *testing.T) { testChannelMutexAdd(t, 1000) }
