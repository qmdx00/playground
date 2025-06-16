package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrencyAdd(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100000 {
				count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Want %d, Got %d\n", 1000000, count)
}

func TestConcurrencyAddWithMutex(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100000 {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Want %d, Got %d\n", 1000000, count)
}
