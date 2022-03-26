package logic

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	pb "example.com/cache-service/z_generated"
	"github.com/go-redis/redis/v8"
)

func SetUserKey(rdb *redis.Client, ctx context.Context, user *pb.User) error {
	name := user.Name
	class := user.Class
	roll := strconv.Itoa(int(user.RollNum))

	var key string = "shivang:" + name + ":" + class + ":" + roll
	value := string(user.Metadata)
	err := SetKey(rdb, ctx, key, value)
	return err
}

func GetUserValue(rbd *redis.Client, ctx context.Context, gur *pb.GetUserRequest) (*pb.User, error) {

	username := gur.Name
	userroll := gur.Roll
	partialkey := "shivang:" + username + ":*:" + strconv.Itoa(int(userroll))

	keys, keyerr := rbd.Keys(ctx, partialkey).Result()
	if keyerr != nil {
		log.Fatalf("err: %v", keyerr)
		return nil, keyerr
	}
	if len(keys) == 0 {
		return nil, errors.New("user not found with  given credentials not found\n")
	}

	key := keys[0]

	userclass := strings.Split(key, ":")[2]

	metadatastring, err := rbd.Get(ctx, key).Result()
	user := &pb.User{
		Name:     username,
		Class:    userclass,
		RollNum:  userroll,
		Metadata: []byte(string(strings.Split(metadatastring, ":")[1])),
	}

	return user, err
}
