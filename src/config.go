package src

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	ServerStartTime time.Time
	Host            string
	Port            int32

	// Service
	Name    string
	Version string

	// DB
	DbUri  string
	DbName string

	// Message Broker
	MessageBrokerUsername     string
	MessageBrokerPassword     string
	MessageBrokerHost         string
	MessageBrokerPort         string
	MessageBrokerExchangeName string

	// Food Service
	FoodServiceExchange              string
	FoodServiceRouteAddFoodChangeLog string
}

func NewAppConfig() (*AppConfig, error) {
	// Get Port Env
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to get env port :%v", err)
	}
	dbUri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	return &AppConfig{
		ServerStartTime: time.Now(),
		Host:            os.Getenv("HOST"),
		Port:            int32(port),

		Name:    os.Getenv("SERVICE_NAME"),
		Version: os.Getenv("SERVICE_VERSION"),

		DbUri:  dbUri,
		DbName: os.Getenv("DB_NAME"),

		MessageBrokerUsername:     os.Getenv("MESSAGE_BROKER_USERNAME"),
		MessageBrokerPassword:     os.Getenv("MESSAGE_BROKER_PASSWORD"),
		MessageBrokerHost:         os.Getenv("MESSAGE_BROKER_HOST"),
		MessageBrokerPort:         os.Getenv("MESSAGE_BROKER_PORT"),
		MessageBrokerExchangeName: os.Getenv("MESSAGE_BROKER_EXCHANGE_NAME"),

		FoodServiceExchange:              os.Getenv("FOOD_SERVICE_EXCHANGE"),
		FoodServiceRouteAddFoodChangeLog: os.Getenv("FOOD_SERVICE_ROUTE_ADD_FOOD_CHANGE_LOG"),
	}, nil
}
