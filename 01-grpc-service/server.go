package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":8080"
)

type Server struct{}

func (s *Server) FizzBuzz(context context.Context, in *FizzBuzzRequest) (*FizzBuzzResponse, error) {
	return &FizzBuzzResponse{Output: "1"}, nil
}

func (s *Server) Stats(context context.Context, in *StatsRequest) (*StatsResponse, error) {
	return &StatsResponse{}, nil
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
