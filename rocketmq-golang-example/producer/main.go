package main

import (
	"context"
	"fmt"
	"rocketmq-golang-example/config"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {
	rlog.SetLogLevel(config.LogLevel)

	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{config.BrokerAddr})),
		producer.WithQueueSelector(producer.NewHashQueueSelector()), // HACK: 注意要实现分区顺序消息，该配置必须手动指定 hash 方式，默认的轮训方式会导致顺序消息不生效
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}

	if err := p.Start(); err != nil {
		panic(err)
	}

	// send message
	now := time.Now().Local()
	for i := range 10 {
		msg := primitive.NewMessage(config.MessageTopic, fmt.Appendf(nil, "%d from producer at %s", i, now.Format(time.DateTime)))

		// NOTE: shardingKey 相同的消息会被路由到同一个 Queue，从而保证消息的分区顺序
		shardingKey := fmt.Sprintf("SK_%d", now.Unix())
		msg.WithShardingKey(shardingKey)
		msg.WithTag(config.MessageTag)

		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[SEND] QueueID=%d Key=%s Body=[%s]\n",
			res.MessageQueue.QueueId,
			shardingKey,
			msg.Body)
	}

	if err := p.Shutdown(); err != nil {
		panic(err)
	}
}
