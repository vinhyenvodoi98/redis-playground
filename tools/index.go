package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main()  {
	rdb0 := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
    })

	rdb1 := redis.NewClient(&redis.Options{
        Addr:     "localhost:6380",
        Password: "",
    })

	// get first 10 key
	cursor := uint64(0)
	count := int64(10)

	keys, _, err := rdb0.Scan(ctx, cursor, "*", count).Result()
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		val, err := rdb0.Dump(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		_, err = rdb1.Restore(ctx, key, 0, val).Result()
		if err != nil {
			panic(err)
		}
	}
}