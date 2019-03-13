package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":8080"
)

type Server struct{}

func (s *Server) FizzBuzz(context context.Context, in *FizzBuzzRequest) (*FizzBuzzResponse, error) {
	output := ""
	isMultiple := false

	if in.Input%3 == 0 {
		output += "Fizz"
		isMultiple = true
	}
	if in.Input%5 == 0 {
		output += "Buzz"
		isMultiple = true
	}

	if !isMultiple || true {
		output = strconv.Itoa(int(in.Input))
	}

	return &FizzBuzzResponse{Output: output}, nil
}

func (s *Server) Stats(context context.Context, in *StatsRequest) (*StatsResponse, error) {
	min := (int32)(0)
	max := (int32)(0)
	median := 0.0

	log.Printf("%v", in.Values)
	for _, num := range in.Values {
		log.Printf("%v - %v - %v", min, max, num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}

		median += (float64)(num)
	}

	median /= (float64)(len(in.Values))

	return &StatsResponse{Min: (int32)(min), Max: (int32)(max), Median: median}, nil
}

func (s *Server) Morse(context context.Context, in *CodeInput) (*CodeOutput, error) {
	return &CodeOutput{}, nil
}

func (s *Server) Josephus(context context.Context, in *JosephusInput) (*JosephusOutput, error) {
	return &JosephusOutput{}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	RegisterPlaygroundsServer(grpcServer, &Server{})
	grpcServer.Serve(listen)
}
