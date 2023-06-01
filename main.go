package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	trippb "github.com/youzhicode/ymcar/proto/gen/go"
	"github.com/youzhicode/ymcar/tripservice"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Printf("Failed to listen %v\n", err)
	}

	s := grpc.NewServer()

	trippb.RegisterTripServiceServer(s, tripservice.TripService)

	log.Fatal(s.Serve(listen))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8888",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("Cannot start grpc gateway %v", err)
	}
	err = http.ListenAndServe(":8887", mux)
	if err != nil {
		log.Fatalf("Cannot listen and start %v", err)
	}
}
