package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/cache-service/z_generated"
)

func RunRawClient() {
	c, err := ConnectToCache()
	if err != nil {
		log.Fatalf("error connecting to cache")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Get(ctx, &pb.Key{
		Key: "asd",
	})
	if err != nil {
		log.Fatalf("error: %s\n", err)
	} else {
		fmt.Printf("response value for get raw is: %s\n", string(res.Value))
	}

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
		log.Fatalf("2 set error %v\n", seterr)
	} else {
		fmt.Printf("3 set res %v\n", setres)
	}

}
