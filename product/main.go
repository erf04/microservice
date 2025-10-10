package main

import (
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"product/config"
	pb "product/proto/productpb"
	"product/repository"
	"product/server"
	"product/service"

	"google.golang.org/grpc"
)

func main() {
	// Connect to MongoDB
    conf:= config.LoadConfig()
	// fmt.Println(conf.MongoDB)
    URI:= fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",conf.MongoUsername,conf.MongoPassword,conf.MongoHost,conf.MongoPort,conf.MongoDB)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	// db := client.("products")
    // fmt.Println(db.ReadPreference().String())
    // log.Default()
	// Setup layers
	repo := repository.NewProductRepository(client.Database(conf.MongoDB))
	svc := service.NewProductService(repo)
	grpcServer := grpc.NewServer()
	productServer := server.NewProductServer(svc)

	// Register gRPC
	pb.RegisterProductServiceServer(grpcServer, productServer)

	// Start server
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Product gRPC server running on 0.0.0.0:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
