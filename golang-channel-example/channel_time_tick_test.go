package main

import (
	"fmt"
	"testing"
	"time"
)

// NOTE: 模拟定时器
func Tick(duration time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)

	go func() {
		for {
			time.Sleep(duration)
			ch <- struct{}{}
		}
	}()

	return ch
}

func testChannelTimeTick(_ *testing.T, delay, timeout time.Duration) {
	after := time.After(timeout)
	ticker := Tick(delay)

	for {
		select {
		case <-ticker:
			fmt.Println("tick")
		case <-after:
			return
		}
	}
}

func TestChannelTimeTick_1(t *testing.T) { testChannelTimeTick(t, time.Millisecond*100, time.Second) }
func TestChannelTimeTick_2(t *testing.T) { testChannelTimeTick(t, time.Millisecond*500, time.Second) }
