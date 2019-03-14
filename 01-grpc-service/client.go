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

	log.Printf("FizzBuzz is: %s", r.Output)

	// values := []int32{5, 10, 2, 11, 20}
	values := []int32{2, 2, 3, 4, 8, 8}
	r2, err := c.Stats(ctx, &StatsRequest{Values: values})

	log.Printf("Stats is: %v, %v, %v", r2.Min, r2.Max, r2.Median)
}
