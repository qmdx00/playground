package config

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	Broker = "127.0.0.1"
	Port   = 1883

	TestTopic = "topic/test"
)

// NOTE: MQTT 发布者指定的 PUBLISH 报文 QoS 等级与订阅者在 SUBSCRIBE 时指定的 QoS 等级不一致时，Broker 会执行 QoS 降级。
// NOTE: 最终实际投递给订阅者的消息 QoS 等级是两者中较低的那个等级。即最终生效的 QoS = min(发布 QoS, 订阅 QoS)。
const (
	QoS0 byte = iota // 最多交付一次
	QoS1             // 至少交付一次
	QoS2             // 只交付一次
)

var OnConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Printf("Connected at %s\n", time.Now().Format(time.DateTime))
}

var ReconnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, opts *mqtt.ClientOptions) {
	fmt.Printf("Reconnected: %v\n", opts)
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
