package main

import (
	"Rusprofile/proto"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"sync"
)

func grpcGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gtw := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterRusprofileWrapperHandlerFromEndpoint(ctx, gtw, addr, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", gtw)
}

func StartHppServe(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial(
		addr, grpc.WithTransportCredentials(
			insecure.
				NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if err = grpcGateway(); err != nil {
		log.Fatal(err)
	}
}
