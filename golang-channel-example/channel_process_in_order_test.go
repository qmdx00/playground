package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// NOTE: 循环顺序打印 cat、dog、fish
func TestProcessInOrder_v1(t *testing.T) {
	var times uint = 5
	var wg sync.WaitGroup

	catCh := make(chan struct{})
	dogCh := make(chan struct{})
	fishCh := make(chan struct{})

	wg.Add(3)

	go func() {
		for range times {
			<-catCh
			fmt.Println("cat")
			time.Sleep(time.Millisecond * 100)
			dogCh <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		for range times {
			<-dogCh
			fmt.Println("dog")
			time.Sleep(time.Millisecond * 100)
			fishCh <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		defer wg.Done()

		for i := range times {
			<-fishCh
			fmt.Println("fish")
			time.Sleep(time.Millisecond * 100)

			// NOTE: 避免 fish 最后一次向 catCh 发送时，因 cat 已退出而阻塞
			// NOTE: wg.Wait() 处永远等待，因为 fish goroutine 未完成，从而形成死锁。
			if i < times-1 {
				catCh <- struct{}{}
			} else {
				break
			}
		}
	}()

	catCh <- struct{}{}

	wg.Wait()
}

func TestProcessInOrder_v2(t *testing.T) {
	var times uint = 5
	var wg sync.WaitGroup

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	wg.Add(3)
	go handle(&wg, "cat", times, ch1, ch2)
	go handle(&wg, "dog", times, ch2, ch3)
	go handle(&wg, "fish", times, ch3, ch1)

	ch1 <- struct{}{}

	wg.Wait()
}

func handle(wg *sync.WaitGroup, name string, times uint, inputCh, outputCh chan struct{}) {
	defer wg.Done()

	for range times {
		<-inputCh
		fmt.Println(name)
		time.Sleep(time.Millisecond * 100)
		outputCh <- struct{}{} // NOTE: 由于使用的是有缓冲的 channel，此处可以安全退出，但是最后一轮的 catCh 中会留下一个未处理的 struct{}
	}
}
