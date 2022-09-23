package main

import (
	"Rusprofile/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func StartGrpcServer(wg *sync.WaitGroup) {
	defer wg.Done()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterRusprofileWrapperServer(server, &RusprofileWrapper{})
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
