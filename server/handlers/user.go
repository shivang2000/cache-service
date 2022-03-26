package handlers

import (
	"context"
	"fmt"

	"example.com/cache-service/server/logic"
	pb "example.com/cache-service/z_generated"
)

func (c *CacheServiceServer) SetUser(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	rdb := c.rdb
	err := logic.SetUserKey(rdb, ctx, user)
	fmt.Printf("rpc method call SetUser with user: %v\n", user)
	return &pb.Empty{}, err
}

func (c *CacheServiceServer) GetUser(ctx context.Context, gur *pb.GetUserRequest) (*pb.User, error) {
	rdb := c.rdb
	user, err := logic.GetUserValue(rdb, ctx, gur)
	fmt.Printf("rpc method call GetUser with user: %v\n", gur)
	return user, err
}
