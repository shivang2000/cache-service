package main

import (
	"fmt"
	"log"
	"net"

	raw "example.com/cache-service/server/handlers"
	pb "example.com/cache-service/z_generated"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

const portNo = ":50000"

func main() {
	list, err := net.Listen("tcp", portNo)
	if err != nil {
		fmt.Printf("error listing on port: %v\n error: %v", portNo, err)
	}

	s := grpc.NewServer()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pb.RegisterCacheServer(s, raw.NewCacheServiceServer(rdb))
	log.Printf("server listening at %v", list.Addr())
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
