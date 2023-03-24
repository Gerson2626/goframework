package goframework

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	ConsumerFunc    func(ctx *ConsumerContext)
	ConsumerContext struct {
		context.Context
		RemainingRetries uint16
		Faulted          bool
		Msg              *kafka.Message
	}
	Consumer interface {
		HandleFn()
	}
	Producer[T interface{}] interface {
		Publish(ctx context.Context, msgs ...*T) error
	}
)

func (cc ConsumerContext) Deadline() (deadline time.Time, ok bool) {
	return cc.Context.Deadline()
}

func (cc ConsumerContext) Done() <-chan struct{} {
	return cc.Context.Done()
}

func (cc ConsumerContext) Err() error {
	return cc.Context.Err()
}

func (cc ConsumerContext) Value(key any) any {
	return cc.Context.Value(key)
}
