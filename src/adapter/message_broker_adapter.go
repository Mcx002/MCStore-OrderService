package adapter

import (
	"context"
	"fmt"
	"mcx002/orderService/src"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueBindInformation struct {
	Name     string
	RouteKey string
}

type MessageBrokerAdapter struct {
	Ch        MessageBrokerChannelInterface
	Conn      MessageBrokerConnectionInterface
	AppConfig *src.AppConfig
	listQueue []QueueBindInformation
}

type MessageBrokerInterface interface {
	Connect(*src.AppConfig) error
	Publish(ctx context.Context, exchange string, routeKey string, body []byte, opt PublishOption) error
	Close()
}

type MessageBrokerConnectionInterface interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

type MessageBrokerChannelInterface interface {
	Close() error
	PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	QueueUnbind(name, key, exchange string, args amqp.Table) error
	QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error)
	ExchangeDelete(name string, ifUnused, noWait bool) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

func (mba *MessageBrokerAdapter) Connect() error {
	appConfig := mba.AppConfig

	// Prepare connection url
	connectionUrl := fmt.Sprintf(
		`amqp://%s:%s@%s:%s/`,
		appConfig.MessageBrokerUsername,
		appConfig.MessageBrokerPassword,
		appConfig.MessageBrokerHost,
		appConfig.MessageBrokerPort,
	)

	// Connect to Message Broker
	conn, err := amqp.Dial(connectionUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to rabbitmq: %v", err)
	}

	// Set the connection to adapter
	mba.Conn = conn

	// Get a channel
	ch, err := conn.Channel()

	// Set the channel to adapter
	mba.Ch = ch

	return err
}

type PublishOption struct {
	Durable      bool
	Persistent   bool
	ExchangeType string
}

func (mba MessageBrokerAdapter) Publish(
	ctx context.Context,
	exchange string,
	routeKey string,
	body []byte,
	opt PublishOption,
) error {
	// prepare default value option
	if opt.ExchangeType == "" {
		opt.ExchangeType = "direct"
	}

	// prepare exchange
	err := mba.Ch.ExchangeDeclare(
		exchange,
		opt.ExchangeType,
		opt.Durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange")
	}

	// set delivery mode
	deliveryMode := amqp.Transient
	if opt.Persistent {
		deliveryMode = amqp.Persistent
	}

	// publish with context
	err = mba.Ch.PublishWithContext(ctx, exchange, routeKey, false, false,
		amqp.Publishing{
			DeliveryMode: deliveryMode,
			ContentType:  "application/json",
			Body:         body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish the message, Exchange: %s, RouteKey: %s, Body: %v", exchange, routeKey, body)
	}

	return nil
}

func (mba *MessageBrokerAdapter) Close() error {
	var err error
	for _, q := range mba.listQueue {
		// Unbind Queue
		err = mba.Ch.QueueUnbind(q.Name, q.RouteKey, mba.AppConfig.MessageBrokerExchangeName, amqp.Table{})
		if err != nil {
			return fmt.Errorf("failed to unbind queue, queue name: %s. err: %v", q.Name, err)
		}

		// Delete Queue
		_, err = mba.Ch.QueueDelete(q.Name, true, true, false)
		if err != nil {
			return fmt.Errorf("failed to delete queue, queue name: %s. err: %v", q.Name, err)
		}
	}

	// Close the channel
	err = mba.Ch.Close()
	if err != nil {
		return fmt.Errorf("failed to close channel message broker: %v", err)
	}

	// Close the connection
	err = mba.Conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection message broker: %v", err)
	}

	return nil
}

type ListenOption struct {
	Durable bool
}

func (mba *MessageBrokerAdapter) ListenTo(queueName string, opt ListenOption, fun func(d amqp.Delivery) error) error {
	exchange := mba.AppConfig.MessageBrokerExchangeName

	// prepare exchange
	err := mba.Ch.ExchangeDeclare(
		exchange,
		"direct",
		opt.Durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange")
	}

	// declare queue
	q, err := mba.Ch.QueueDeclare(
		queueName,
		opt.Durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	// bind queue
	exchangeName := mba.AppConfig.MessageBrokerExchangeName
	routeKey := fmt.Sprintf("%s-%s", exchangeName, q.Name)
	err = mba.Ch.QueueBind(
		q.Name,
		routeKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to binding the queue: %v", err)
	}

	// Consume messages
	msgs, err := mba.Ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume the messages: %v", err)
	}

	// listen to messages
	go func() {
		for d := range msgs {
			err = fun(d)
			if err != nil {
				d.Nack(false, false)
			}
			d.Ack(false)
		}
	}()

	// Add the queue to list queue
	mba.listQueue = append(mba.listQueue, QueueBindInformation{
		Name:     q.Name,
		RouteKey: routeKey,
	})
	return nil
}
