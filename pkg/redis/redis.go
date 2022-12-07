package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func NewClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
