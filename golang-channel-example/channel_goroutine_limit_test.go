package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

// NOTE: 控制并发的 Goroutine 数量
func testChannelGoroutineLimit(_ *testing.T, limit, total uint) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, limit)

	for i := range total {
		ch <- struct{}{}

		wg.Add(1)
		go func(i uint) {
			defer wg.Done()

			log.Println(i)
			time.Sleep(time.Second)

			<-ch
		}(i)
	}

	wg.Wait()
}

func TestChannelGoroutineLimit_1_2(t *testing.T)   { testChannelGoroutineLimit(t, 1, 2) }
func TestChannelGoroutineLimit_10_20(t *testing.T) { testChannelGoroutineLimit(t, 10, 20) }
func TestChannelGoroutineLimit_20_20(t *testing.T) { testChannelGoroutineLimit(t, 20, 20) }
