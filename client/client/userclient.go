package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/cache-service/z_generated"
)

func RunUserCLient() {

	c, err := ConnectToCache()
	if err != nil {
		log.Fatalf("error connecting to cache\n")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, seterr := c.SetUser(ctx, &pb.User{
		Name:     "Alice",
		Class:    "IV",
		RollNum:  15,
		Metadata: []byte("secret message alice is wonderfull"),
	})
	if seterr != nil {
		fmt.Printf("error getting data for %s\n", err)
	}
	fmt.Printf("1 response value\n")

	user, geterr := c.GetUser(ctx, &pb.GetUserRequest{
		Name: "Alice",
		Roll: 15,
	})
	if geterr != nil {
		fmt.Printf("error get user data: %v\n", geterr)
	}
	if user != nil {
		fmt.Println("get response: ")
		fmt.Printf("\tName: %v\n", user.Name)
		fmt.Printf("\tclass: %v\n", user.Class)
		fmt.Printf("\tRoll No: %v\n", user.RollNum)
		fmt.Printf("\tmetadata: %v\n\n", string(user.Metadata))
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

	//setres, seterr := c.Set(ctx, &pb.SetRequest{
	// 	Key:   &pb.Key{Key: "asd"},
	//	Value: &pb.Value{Value: []byte("asda")},
	//})
	//if seterr != nil {
	//	fmt.Printf("2 set error %v\n", seterr)
	//}
	//fmt.Printf("3 set res %v\n", setres)

}
