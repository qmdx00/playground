package main

import (
	"fmt"
	"testing"
	"time"
)

func CallFunc(fn func()) <-chan struct{} {
	ch := make(chan struct{}, 1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		fn()
		ch <- struct{}{}
	}()

	return ch
}

// NOTE: 函数执行超时控制
func testCallFunc(_ *testing.T, cost, timeout time.Duration) {
	fn := func() {
		fmt.Println("doing something")
		time.Sleep(cost)
	}

	select {
	case <-CallFunc(fn):
		fmt.Println("done")
	case <-time.After(timeout):
		fmt.Println("timeout")
	}
}

func TestCallFuncDone(t *testing.T)    { testCallFunc(t, time.Second, time.Second*2) }
func TestCallFuncTimeout(t *testing.T) { testCallFunc(t, time.Second*2, time.Second) }
