package service

import (
	"context"
	"log"
	"mcx002/orderService/src"
	"mcx002/orderService/src/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func NewAppConfigTest() (*src.AppConfig, error) {
	// load .env file
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatalf("cannot load env :%v", err)
	}

	// load env
	return src.NewAppConfig()
}

func TestSetContext(t *testing.T) {
	// get config
	appConfig, err := NewAppConfigTest()
	assert.Nil(t, err)

	// create service
	service := &Service{
		Config: appConfig,
		Repo:   &RepositoryMock{},
	}

	// validate context
	ctx := context.Background()
	service.SetContext(ctx)

	assert.Equal(t, ctx, service.ctx)
}

func NewServiceTest(t *testing.T) *Service {
	// get config
	appConfig, err := NewAppConfigTest()
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	// create service
	service := &Service{
		Config: appConfig,
		Repo:   &RepositoryMock{},
	}

	// create context
	ctx := context.Background()
	service.SetContext(ctx)

	return service
}

var (
	MockGetOrderCount func() (int64, error)
	MockInsertOrder   func(*model.Order, []model.OrderItem) (string, error)
)

// prepare repository mock
type RepositoryMock struct{}

func (r RepositoryMock) SetContext(context.Context) {}
func (r RepositoryMock) InsertOrder(order *model.Order, items []model.OrderItem) (string, error) {
	return MockInsertOrder(order, items)
}
func (r RepositoryMock) GetOrderCountToday() (int64, error) {
	return MockGetOrderCount()
}
