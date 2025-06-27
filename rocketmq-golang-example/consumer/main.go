package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rocketmq-golang-example/config"
	"syscall"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {
	rlog.SetLogLevel(config.LogLevel)

	options := []consumer.Option{
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{config.BrokerAddr})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumerOrder(true), // NOTE: 开启顺序消费
	}

	pushConsumer(options)
}

func pushConsumer(options []consumer.Option) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	c, err := rocketmq.NewPushConsumer(append(options, consumer.WithGroupName("GID_TEST_1"))...)
	if err != nil {
		panic(err)
	}

	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: config.MessageTag,
	}

	messageHandle := func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("[RECV] QueueID=%d Offset=%d Body=[%s]\n",
				msg.Queue.QueueId,
				msg.QueueOffset,
				msg.Body)
		}
		return consumer.ConsumeSuccess, nil
	}

	if err := c.Subscribe(config.MessageTopic, selector, messageHandle); err != nil {
		panic(err)
	}

	if err := c.Start(); err != nil {
		panic(err)
	}

	<-sigs
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
}
