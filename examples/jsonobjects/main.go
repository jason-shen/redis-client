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

	sample2 := map[string]interface{} {
		"firstname": "hello2",
		"lastname": "world1",
	}

	redis.Create(key, sample)
	data, err := redis.Read(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", data)
	redis.Update(key, sample2)
	data2, err := redis.Read(key)
	fmt.Println(data2)
}