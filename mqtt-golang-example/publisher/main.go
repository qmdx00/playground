package main

import (
	"fmt"
	"mqtt-golang-example/config"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Broker, config.Port))
	opts.SetClientID("publisher-test")
	opts.SetCleanSession(false)

	// event handler
	opts.OnConnect = config.OnConnectHandler
	opts.OnReconnecting = config.ReconnectHandler
	opts.OnConnectionLost = config.ConnectLostHandler

	// client connect
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// publish some messages
	for i := range 10 {
		payload := fmt.Sprintf("Message %d at %s", i, time.Now().Format(time.DateTime))
		if token := client.Publish(config.TestTopic, config.QoS2, false, payload); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		time.Sleep(time.Millisecond * 100)
		fmt.Println("Published message:", payload)
	}

	// publish end, disconnect
	client.Disconnect(250)
}
