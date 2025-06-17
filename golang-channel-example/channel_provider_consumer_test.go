package main

import (
	"fmt"
	"testing"
	"time"
)

// 向 channel 里投递数据
func Provider(ch chan<- int, nums ...int) {
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

// 从 channel 里消费数据
func Consumer(ch <-chan int, done chan<- struct{}, fn func(int)) {
	for num := range ch {
		fn(num)
	}
	done <- struct{}{}
	close(done)
}

// NOTE: 生产者-消费者数据通道
func testProviderAndConsumer(_ *testing.T, consumeFn func(int), nums ...int) {
	// 数据交换 channel
	ch := make(chan int, 10)
	// 控制 channel 退出
	done := make(chan struct{})

	go Provider(ch, nums...)
	go Consumer(ch, done, consumeFn)

	// 接收退出并进行超时处理
	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	}
}

func ConsumeFunc(num int) {
	fmt.Println(num)
	time.Sleep(time.Millisecond * 100)
}

func TestProviderAndConsumer_1(t *testing.T) { testProviderAndConsumer(t, ConsumeFunc, 1, 2, 3, 4, 5) }
func TestProviderAndConsumer_2(t *testing.T) {
	testProviderAndConsumer(t, ConsumeFunc, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
}
