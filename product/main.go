package main

import (
	"fmt"
	"log"

	// "net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"product/config"
	"product/handlers"
	natsclient "product/nats"

	// pb "product/proto/productpb"
	"product/repository"
	"product/service"
)

func main() {
	// Connect to MongoDB
    conf:= config.LoadConfig()
    URI:= fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",conf.MongoUsername,conf.MongoPassword,conf.MongoHost,conf.MongoPort,conf.MongoDB)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewProductRepository(client.Database(conf.MongoDB))
	svc := service.NewProductService(repo)
	// register NATS
	productHandler := handlers.New(svc)
	natsclient := natsclient.New()
	err = natsclient.Connect(fmt.Sprintf("%s:%s",conf.NatsServer,conf.NatsPort))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🚀 Nats server running successfully")
	// register events
	natsclient.RegisterRPCHandler("product.getbyid", "product-workers", productHandler.GetProductByID)
	// grpcServer := grpc.NewServer()
	// productServer := server.NewProductServer(svc)
	select {}
	// Register gRPC
	// pb.RegisterProductServiceServer(grpcServer, productServer)

	// Start server
	// _, err = net.Listen("tcp", "0.0.0.0:50051")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := grpcServer.Serve(listener); err != nil {
	// 	log.Fatal(err)
	// }
}
