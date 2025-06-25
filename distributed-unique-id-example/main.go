package main

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake/v2"
)

func main() {
	// pool := NewIDPoolWithCond(10, sonyflake.Settings{})
	pool := NewIDPoolWithChannel(10, sonyflake.Settings{})

	for i := range 3 {
		go func(gid int) {
			for {
				fmt.Printf("[gid %d] Get ID: %d\n", gid, pool.GetID())
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}

	pool.Run()
}
