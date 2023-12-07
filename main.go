package main

import (
	"context"
	"fmt"
	"log"
	"mcx002/orderService/src"
	"mcx002/orderService/src/adapter"
	"mcx002/orderService/src/model"
	"mcx002/orderService/src/proto_gen"
	"mcx002/orderService/src/repository"
	"mcx002/orderService/src/server"
	"mcx002/orderService/src/service"
	"net"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func loadEnv() (*src.AppConfig, error) {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("cannot load env :%v", err)
	}

	// Load Env Data
	return src.NewAppConfig()
}

func getDBCollections(ctx context.Context, appConfig *src.AppConfig) (*model.CollectionList, *mongo.Client, error) {
	dbAdapter := &adapter.MongoDbAdapter{}
	err := dbAdapter.Connect(ctx, appConfig.DbUri)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to db :%v", err)
	}
	client := dbAdapter.GetClient()
	db := client.Database(appConfig.DbName)

	// Init Collection
	return model.NewCollectionList(db), client, nil
}

func prepareServer(appConfig *src.AppConfig) (*grpc.Server, *adapter.MessageBrokerAdapter, error) {
	// create program context
	ctx := context.Background()

	// get collections
	coll, client, err := getDBCollections(ctx, appConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get collections :%v", err)
	}

	// Init Repository
	repo := &repository.Repository{
		Collections: coll,
		Client:      client,
		Config:      appConfig,
	}
	repo.SetContext(ctx)

	// Init Service
	service := &service.Service{
		Config: appConfig,
		Repo:   repo,
	}

	svr := &server.Server{
		Config:  appConfig,
		Service: service,
	}

	mba, err := svr.ListenToMessageBroker()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen to message broker: %v", err)
	}

	// set repo broker
	repo.SetBroker(mba)

	// Register the server
	s := grpc.NewServer()
	proto_gen.RegisterOrderServer(s, svr)

	// Init server
	return s, mba, nil
}

func main() {
	// get configuration
	appConfig, err := loadEnv()
	if err != nil {
		log.Fatalf("failed to load appConfig :%v", err)
	}

	// prepare grpc server
	s, mba, err := prepareServer(appConfig)
	defer mba.Close()
	if err != nil {
		log.Fatalf("failed to prepare server :%v", err)
	}

	// Listen to TCP protocol connection
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	// Run Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
