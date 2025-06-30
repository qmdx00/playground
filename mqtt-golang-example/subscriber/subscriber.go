package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client mqtt.Client
	quit   chan os.Signal

	channel chan mqtt.Message
	handle  func(msg mqtt.Message)
}

func NewSubscriber(opts *mqtt.ClientOptions, handle func(msg mqtt.Message)) *Subscriber {
	return &Subscriber{
		client:  mqtt.NewClient(opts),
		quit:    make(chan os.Signal),
		channel: make(chan mqtt.Message),
		handle:  handle,
	}
}

func (s *Subscriber) Subscribe(topic string, qos byte) {
	signal.Notify(s.quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
		panic("connect failed")
	}

	s.client.Subscribe(topic, qos, func(_ mqtt.Client, msg mqtt.Message) {
		s.channel <- msg
	})

	s.readLoop()
}

func (s *Subscriber) readLoop() {
	defer s.client.Disconnect(250)
	for {
		select {
		case msg := <-s.channel:
			s.handle(msg)
		case <-s.quit:
			fmt.Println("quit subscribe")
			return
		}
	}
}
