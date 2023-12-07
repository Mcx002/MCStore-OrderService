package repository

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBrokerMockConnection struct{}
type MessageBrokerMockChannel struct{}

var (
	MessageBrokerMockConnectionChannel      func() (*amqp.Channel, error)
	MessageBrokerMockConnectionClose        func() error
	MessageBrokerMockChannelClose           func() error
	MessageBrokerMockChannelPublish         func(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	MessageBrokerMockChannelExchangeDeclare func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	MessageBrokerMockChannelQueueDeclare    func(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	MessageBrokerMockChannelQueueBind       func(name, key, exchange string, noWait bool, args amqp.Table) error
	MessageBrokerMockChannelQueueUnbind     func(name, key, exchange string, args amqp.Table) error
	MessageBrokerMockChannelQueueDelete     func(name string, ifUnused, ifEmpty, noWait bool) (int, error)
	MessageBrokerMockChannelExchangeDelete  func(name string, ifUnused, noWait bool) error
	MessageBrokerMockChannelConsume         func(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
)

func (m MessageBrokerMockConnection) Channel() (*amqp.Channel, error) {
	return MessageBrokerMockConnectionChannel()
}

func (m MessageBrokerMockConnection) Close() error {
	return MessageBrokerMockConnectionClose()
}

func (mbdc MessageBrokerMockChannel) Close() error {
	return MessageBrokerMockChannelClose()
}

func (mbdc MessageBrokerMockChannel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return MessageBrokerMockChannelPublish(ctx, exchange, key, mandatory, immediate, msg)
}

func (mbdc MessageBrokerMockChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return MessageBrokerMockChannelExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (mbdc MessageBrokerMockChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return MessageBrokerMockChannelQueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (mbdc *MessageBrokerMockChannel) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	return MessageBrokerMockChannelQueueBind(name, key, exchange, noWait, args)
}

func (mbdc *MessageBrokerMockChannel) QueueUnbind(name, key, exchange string, args amqp.Table) error {
	return MessageBrokerMockChannelQueueUnbind(name, key, exchange, args)
}

func (mbdc *MessageBrokerMockChannel) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	return MessageBrokerMockChannelQueueDelete(name, ifUnused, ifEmpty, noWait)
}

func (mbdc *MessageBrokerMockChannel) ExchangeDelete(name string, ifUnused, noWait bool) error {
	return MessageBrokerMockChannelExchangeDelete(name, ifUnused, noWait)
}

func (mbdc *MessageBrokerMockChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return MessageBrokerMockChannelConsume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}
