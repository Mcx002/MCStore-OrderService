package service

import (
	"context"
	"fmt"
	"mcx002/orderService/src"
	"mcx002/orderService/src/model"
	"mcx002/orderService/src/proto_gen"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init(t *testing.T) (context.Context, *src.AppConfig, *Service) {
	ctx := context.Background()

	// get env test
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Errorf("cannot load env :%v", err)
	}

	// load config
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Errorf("failed to load configuration :%v", err)
	}

	service := &Service{
		ctx:    ctx,
		Config: appConfig,
		Repo:   &RepositoryMock{},
	}

	return ctx, appConfig, service
}

func TestCreateOrder(t *testing.T) {
	_, _, service := Init(t)

	t.Run("Should failed to get order count", func(t *testing.T) {
		// mock get order count
		MockGetOrderCount = func() (int64, error) {
			return 0, fmt.Errorf("error")
		}

		// execute create order
		result, err := service.CreateOrder(&proto_gen.OrderDto{})

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "failed to get order count")
	})

	t.Run("Should failed to insert order", func(t *testing.T) {
		// mock get order count
		MockGetOrderCount = func() (int64, error) {
			return 2, nil
		}

		// mock insert order
		MockInsertOrder = func(o *model.Order, oi []model.OrderItem) (string, error) {
			return "", fmt.Errorf("error")
		}

		// execute create order
		result, err := service.CreateOrder(&proto_gen.OrderDto{
			Items: []*proto_gen.OrderDto_Item{
				{
					Qty: 5,
				},
			},
			Invoice: &proto_gen.InvoiceDto{},
		})

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "failed to insert order")
	})

	t.Run("Should success create order", func(t *testing.T) {
		// mock get order count
		MockGetOrderCount = func() (int64, error) {
			return 2, nil
		}

		// mock insert order
		MockInsertOrder = func(o *model.Order, oi []model.OrderItem) (string, error) {
			return "id", nil
		}

		// execute create order
		result, err := service.CreateOrder(&proto_gen.OrderDto{
			Items: []*proto_gen.OrderDto_Item{
				{
					Qty: 5,
				},
			},
			Invoice: &proto_gen.InvoiceDto{},
		})

		assert.NotNil(t, result, "")
		assert.Nil(t, err)
		assert.Equal(t, result.Version, int32(1))
	})
}
