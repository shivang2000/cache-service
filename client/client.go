package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/cache-service/z_generated"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()

	c := pb.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Get(ctx, &pb.Key{
		Key: "asd",
	})
	if err != nil {
		fmt.Printf("error getting data for %s\n", err)
	}
	fmt.Printf("1 response value %s\n", string(res.Value))

	// setrequest := pb.SetRequest{
	// 	Key:   &pb.Key{Key: "asd"},
	// 	Value: &pb.Value{Value: []byte{"asd"}},
	// }
	// res, err := c.Set(ctx, setrequest)

	// if err != nil {
	// 	fmt.Printf("error getting data for %s", err)
	// }
	// fmt.Printf("response value %s", res)

	setres, seterr := c.Set(ctx, &pb.SetRequest{
		Key:   &pb.Key{Key: "asd"},
		Value: &pb.Value{Value: []byte("asda")},
	})
	if seterr != nil {
		fmt.Printf("2 set error %v\n", seterr)
	}
	fmt.Printf("3 set res %v\n", setres)

}
