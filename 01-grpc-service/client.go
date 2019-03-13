package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	defer conn.Close()
	c := NewPlaygroundsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FizzBuzz(ctx, &FizzBuzzRequest{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greetings: %s", r.Output)
}
