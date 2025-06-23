package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Queue[T any] struct {
	items []T
	cap   uint
	cond  *sync.Cond
}

func NewQueue[T any](cap uint) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, cap),
		cap:   cap,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for uint(len(q.items)) >= q.cap {
		fmt.Println("Queue is full, enqueue waiting...")
		q.cond.Wait()
	}

	q.items = append(q.items, item)
	fmt.Printf("Enqueued: %v, Queue items: %v\n", item, q.items)

	q.cond.Broadcast()
}

func (q *Queue[T]) Dequeue() T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.items) == 0 {
		fmt.Println("Queue is empty, dequeue waiting...")
		q.cond.Wait()
	}

	item := q.items[0]
	q.items = q.items[1:]
	fmt.Printf("Dequeued: %v, Queue items: %v\n", item, q.items)

	q.cond.Broadcast()
	return item
}

func TestQueueWithCondBroadcast(t *testing.T) {
	queue := NewQueue[int](3)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// consumer
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Millisecond * 500)
			queue.Dequeue()
		}
	}()

	// producer
	go func() {
		defer wg.Done()
		for i := range 20 {
			time.Sleep(time.Millisecond * 100)
			queue.Enqueue(i)
		}
	}()

	wg.Wait()
}

func TestBroadcastWithChannel(t *testing.T) {
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				<-ch
				fmt.Printf("channel %d received\n", i)
				break
			}
		}()
	}

	time.Sleep(time.Second)
	close(ch)
	fmt.Println("close channel")

	wg.Wait()
}
