package gRPC

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var Connection = func() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	clientConnection, err := grpc.Dial(os.Getenv("MICROSERVICE_BOOKS_GRPC_ADDRESS"), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return clientConnection
}
