package main

import (
	"sync"

	"github.com/sony/sonyflake/v2"
)

type IDPoolWithCond struct {
	ids  map[int64]struct{}
	cap  uint
	cond *sync.Cond

	sf *sonyflake.Sonyflake
}

func NewIDPoolWithCond(cap uint, settings ...sonyflake.Settings) *IDPoolWithCond {
	if cap == 0 {
		panic("ID pool cap must be greater than 0")
	}

	st := sonyflake.Settings{}
	if len(settings) > 0 {
		st = settings[0]
	}

	sf, err := sonyflake.New(st)
	if err != nil {
		panic("Failed to create Sonyflake instance: " + err.Error())
	}

	return &IDPoolWithCond{
		ids:  make(map[int64]struct{}, cap),
		cap:  cap,
		cond: sync.NewCond(&sync.Mutex{}),
		sf:   sf,
	}
}

func (p *IDPoolWithCond) Run() {
	for {
		p.cond.L.Lock()

		// Check if the pool is full
		for uint(len(p.ids)) >= p.cap {
			p.cond.Wait()
		}

		// Generate a new ID
		id, err := p.sf.NextID()
		if err != nil {
			continue
		}
		p.ids[id] = struct{}{}

		// Notify any waiting goroutines that an ID has been added
		p.cond.Broadcast()
		p.cond.L.Unlock()
	}
}

func (p *IDPoolWithCond) GetID() int64 {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	// Check if the pool is empty
	for len(p.ids) <= 0 {
		p.cond.Wait()
	}

	// Get the first ID from the map
	var id int64
	for id = range p.ids {
		break
	}

	// Remove the ID from the pool
	delete(p.ids, id)

	p.cond.Broadcast()
	return id
}
