package main

import (
	"context"
	"rocketmq-golang-example/config"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	// NOTE: 初始化创建测试用的 TOPIC
	topicAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{config.BrokerAddr})))
	if err != nil {
		panic(err)
	}

	if err = topicAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(config.MessageTopic),
		admin.WithBrokerAddrCreate("127.0.0.1:10911"),
	); err != nil {
		panic(err)
	}
}
