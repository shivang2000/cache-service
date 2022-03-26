package handlers

import (
	pb "example.com/cache-service/z_generated"
	"github.com/go-redis/redis/v8"
)

const (
	port = ":50000"
)

func NewCacheServiceServer(rdb *redis.Client) (c *CacheServiceServer) {
	if rdb != nil {
		return &CacheServiceServer{rdb: rdb}
	}
	return &CacheServiceServer{}
}

type CacheServiceServer struct {
	pb.UnimplementedCacheServer
	rdb *redis.Client
}
