package client

import (
	pb "example.com/cache-service/z_generated"
	"google.golang.org/grpc"
)

const Address string = "localhost:50000"

func ConnectToCache() (pb.CacheClient, error) {
	connectToGRPC, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	conn := pb.NewCacheClient(connectToGRPC)

	return conn, nil
}
