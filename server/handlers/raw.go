package handlers

import (
	"context"
	"errors"
	"fmt"

	pb "example.com/cache-service/z_generated"
)

func (c *CacheServiceServer) Set(ctx context.Context, sr *pb.SetRequest) (*pb.Empty, error) {
	fmt.Printf("rpc Set method called with get following key and value: %v %v  respectively\n", sr.GetKey(), sr.GetValue())
	return &pb.Empty{}, errors.New("not implemented yet. <shivang> will implement me")
}

func (c *CacheServiceServer) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	var s string = "shivang will implement me"

	fmt.Printf("rpc Get method called with key: %v\n", key)

	return &pb.Value{
		Value: []byte(s),
	}, nil
}
