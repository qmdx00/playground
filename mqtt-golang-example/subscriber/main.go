package main

import (
	"fmt"
	"mqtt-golang-example/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Broker, config.Port))
	opts.SetClientID("subscriber-test")
	opts.SetCleanSession(false)

	// event handler
	opts.OnConnect = config.OnConnectHandler
	opts.OnReconnecting = config.ReconnectHandler
	opts.OnConnectionLost = config.ConnectLostHandler

	var messageHandle = func(msg mqtt.Message) {
		fmt.Printf("[QoS=%d, messageID=%d, topic=%s] Received message: [%s]\n",
			msg.Qos(),
			msg.MessageID(),
			msg.Topic(),
			msg.Payload())
	}

	// subscribe
	subscriber := NewSubscriber(opts, messageHandle)
	subscriber.Subscribe(config.TestTopic, config.QoS2)
}

// ================================== origin usage ==============================

// func main() {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

// 	opts := mqtt.NewClientOptions()
// 	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Broker, config.Port))
// 	opts.SetClientID("subscriber-test")
// 	opts.SetCleanSession(false)

// 	// message handler
// 	var messageHandler mqtt.MessageHandler = func(_ mqtt.Client, msg mqtt.Message) {
// 		fmt.Printf("[QoS=%d, messageID=%d, topic=%s] Received message: [%s]\n",
// 			msg.Qos(),
// 			msg.MessageID(),
// 			msg.Topic(),
// 			msg.Payload())
// 	}

// 	// event handler
// 	opts.OnConnect = config.OnConnectHandler
// 	opts.OnReconnecting = config.ReconnectHandler
// 	opts.OnConnectionLost = config.ConnectLostHandler

// 	// client connect and subscribe
// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		panic(token.Error())
// 	}
// 	if token := client.Subscribe(config.TestTopic, config.QoS2, messageHandler); token.Wait() && token.Error() != nil {
// 		panic(token.Error())
// 	}
// 	fmt.Printf("Subscribed to topic: %s\n", config.TestTopic)

// 	// wait exit
// 	<-sigs
// 	client.Disconnect(250)
// }
