package adapter

import (
	"context"
	"fmt"
	"mcx002/orderService/src"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func Init(t *testing.T) *src.AppConfig {
	// load .env file
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Errorf("cannot load env :%v", err)
	}

	// get configuration
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Errorf("cannot get configuration :%v", err)
	}

	return appConfig
}

func TestMessageBrokerConnect(t *testing.T) {
	appConfig := Init(t)

	// prepare message broker
	mb := &MessageBrokerAdapter{
		AppConfig: appConfig,
	}

	t.Run("Should failed to connect", func(t *testing.T) {
		// set mock message broker host
		appConfig.MessageBrokerHost = "MBHost"

		// execute message broker connect
		err := mb.Connect()

		assert.Contains(t, err.Error(), "failed to connect to rabbitmq")
	})
	t.Run("Should connected", func(t *testing.T) {
		// set message broker host
		appConfig.MessageBrokerHost = "localhost"

		// execute message broker connect
		err := mb.Connect()

		assert.Nil(t, err)
	})
}

func TestMessageBrokerClose(t *testing.T) {
	appConfig := Init(t)

	// prepare message broker adapter
	mb := &MessageBrokerAdapter{
		Ch:        &MessageBrokerMockChannel{},
		Conn:      &MessageBrokerMockConnection{},
		AppConfig: appConfig,
	}

	t.Run("Should failed to unbind queue", func(t *testing.T) {
		// prepare list queue
		mb.listQueue = []QueueBindInformation{
			{
				Name:     "test",
				RouteKey: "test-route",
			},
		}

		// mock unbind queue function
		MessageBrokerMockChannelQueueUnbind = func(name, key, exchange string, args amqp091.Table) error {
			return fmt.Errorf("error")
		}

		// execute message broker close
		err := mb.Close()

		assert.Contains(t, err.Error(), "failed to unbind queue")
	})
	t.Run("Should failed to delete queue", func(t *testing.T) {
		// prepare list queue
		mb.listQueue = []QueueBindInformation{
			{
				Name:     "test",
				RouteKey: "test-route",
			},
		}

		// mock unbind queue function
		MessageBrokerMockChannelQueueUnbind = func(name, key, exchange string, args amqp091.Table) error {
			return nil
		}

		// mock delete queue function
		MessageBrokerMockChannelQueueDelete = func(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
			return 0, fmt.Errorf("error")
		}

		// execute message broker close
		err := mb.Close()

		assert.Contains(t, err.Error(), "failed to delete queue")
	})

	// set list queue empty
	mb.listQueue = []QueueBindInformation{}

	t.Run("Should failed to close message broker channel", func(t *testing.T) {
		// mock close channel
		MessageBrokerMockChannelClose = func() error {
			return fmt.Errorf("Error")
		}

		// execute message broker close
		err := mb.Close()

		assert.Contains(t, err.Error(), "failed to close channel message broker")
	})

	t.Run("Should failed to close message broker connection", func(t *testing.T) {
		// mock close channel
		MessageBrokerMockChannelClose = func() error {
			return nil
		}
		// mock close connection
		MessageBrokerMockConnectionClose = func() error {
			return fmt.Errorf("Error")
		}

		// execute message broker close
		err := mb.Close()

		assert.Contains(t, err.Error(), "failed to close connection message broker")
	})

	t.Run("Should success close connection", func(t *testing.T) {
		// mock close channel
		MessageBrokerMockChannelClose = func() error {
			return nil
		}
		// mock close connection
		MessageBrokerMockConnectionClose = func() error {
			return nil
		}

		// execute message broker close
		err := mb.Close()

		assert.Nil(t, err)
	})
}

func TestMessageBrokerPublish(t *testing.T) {
	mb := &MessageBrokerAdapter{
		Ch:   &MessageBrokerMockChannel{},
		Conn: &MessageBrokerMockConnection{},
	}

	t.Run("Should failed to declare an exchange", func(t *testing.T) {
		// mock exchange declare
		MessageBrokerMockChannelExchangeDeclare = func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
			return fmt.Errorf("error")
		}

		// execute publish
		err := mb.Publish(context.Background(), "testExchange", "testRouteKey", []byte("test"), PublishOption{})

		assert.Contains(t, err.Error(), "failed to declare an exchange")
	})
	t.Run("Should have deliveryModel persistent", func(t *testing.T) {
		// mock exchange declare
		MessageBrokerMockChannelExchangeDeclare = func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
			return nil
		}

		// mock channel publish
		MessageBrokerMockChannelPublish = func(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) error {
			return fmt.Errorf("error")
		}

		// execute publish
		err := mb.Publish(context.Background(), "testExchange", "testRouteKey", []byte("test"), PublishOption{
			Persistent: true,
		})

		assert.Contains(t, err.Error(), "failed to publish the message")
	})
	t.Run("Should failed to publish the message", func(t *testing.T) {
		// mock exchange declare
		MessageBrokerMockChannelExchangeDeclare = func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
			return nil
		}

		// mock channel publish
		MessageBrokerMockChannelPublish = func(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) error {
			return fmt.Errorf("error")
		}

		// execute publish
		err := mb.Publish(context.Background(), "testExchange", "testRouteKey", []byte("test"), PublishOption{})

		assert.Contains(t, err.Error(), "failed to publish the message")
	})
	t.Run("Should succeed publish the message", func(t *testing.T) {
		// mock exchange declare
		MessageBrokerMockChannelExchangeDeclare = func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
			return nil
		}

		// mock channel publish
		MessageBrokerMockChannelPublish = func(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) error {
			return nil
		}

		// execute publish
		err := mb.Publish(context.Background(), "testExchange", "testRouteKey", []byte("test"), PublishOption{
			ExchangeType: "topic",
		})

		assert.Nil(t, err)
	})
}
