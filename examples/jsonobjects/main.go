package main

import (
	"fmt"
	"github.com/jason-shen/redis-client/pkg/client"
)

func main()  {
	redis := client.NewRedisCache("localhost:6379", "", 1, 1000)
	key := "myfirst_key"
	sample := map[string]interface{} {
		"firstname": "hello",
		"lastname": "world",
	}

	redis.Set(key, sample)
	data, err := redis.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", data)
}