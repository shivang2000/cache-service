package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "example.com/cache-service/z_generated"
	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

type CacheServiceServer struct {
	pb.UnimplementedCacheServer
}

func (c *CacheServiceServer) Set(ctx context.Context, sr *pb.SetRequest) (*pb.Empty, error) {
	fmt.Printf("request for set %v %v ", sr.GetKey(), sr.GetValue())
	return &pb.Empty{}, errors.New("not implemented yet. <shivang> will implement me")
}

func (c *CacheServiceServer) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	var s string = "shivang will implement me"
	return &pb.Value{
		Value: []byte(s),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port in use")
	}

	s := grpc.NewServer()
	pb.RegisterCacheServer(s, &CacheServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
