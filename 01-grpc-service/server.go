package main

import (
	"context"
	"fmt"
	"net"
	"sort"
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
	median := (float64)(0)

	// Reference https://forum.golangbridge.org/t/how-to-sort-int32/10529/2
	sort.Slice(in.Values, func(i int, j int) bool {
		return in.Values[i] < in.Values[j]
	})

	if len(in.Values) > 0 {
		min = in.Values[0]
		max = in.Values[len(in.Values)-1]
		median = (float64)(0)

		if len(in.Values)%2 == 0 {
			first := in.Values[(len(in.Values)-1)/2]
			second := in.Values[(len(in.Values)-1)/2+1]
			median = ((float64)(first + second)) / 2
		} else {
			median = (float64)(in.Values[len(in.Values)/2])
		}
	}

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
