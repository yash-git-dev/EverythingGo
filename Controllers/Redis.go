package Controllers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisUse() interface{} {
	ctx := context.Background()
	fmt.Println("yash")
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "123",
			DB:       0,
		},
	)

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error pinging Redis server:", err)
		return err
	}
	fmt.Println("Redis server is running:", pong)

	abc, err := rdb.Set(ctx, "key1", "value1", 0).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(abc)

	val, err := rdb.Get(context.Background(), "key1").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return err
	}
	fmt.Println("Value for key key1:", val)

	return ""
}
