package main

import (
	"context"
	"sync"

	"github.com/sony/sonyflake/v2"
)

type IDPoolWithChannel struct {
	ids chan int64
	cap uint

	sf *sonyflake.Sonyflake

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func NewIDPoolWithChannel(cap uint, settings ...sonyflake.Settings) *IDPoolWithChannel {
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

	ctx, cancel := context.WithCancel(context.Background())
	return &IDPoolWithChannel{
		ids:    make(chan int64, cap),
		cap:    cap,
		sf:     sf,
		ctx:    ctx,
		cancel: cancel,
		wg:     &sync.WaitGroup{},
	}
}

func (p *IDPoolWithChannel) Run() {
	p.wg.Add(1)
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			return
		default:
			// Generate a new ID
			id, err := p.sf.NextID()
			if err != nil {
				continue
			}

			select {
			case p.ids <- id:
				// Send the ID to the channel
			case <-p.ctx.Done():
				return
			}
		}
	}
}

func (p *IDPoolWithChannel) GetID() int64 {
	id := <-p.ids
	return id
}

func (p *IDPoolWithChannel) Close() {
	p.cancel()
	p.wg.Wait()
	close(p.ids)
}
