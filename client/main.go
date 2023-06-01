package main

import (
	"context"
	"fmt"
	"log"

	trippb "github.com/youzhicode/ymcar/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Cannot connect server %v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)
	resp, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})
	if err != nil {
		log.Fatal("Cannot call Gettrip: %v", err)
	}
	fmt.Println(resp)
}
