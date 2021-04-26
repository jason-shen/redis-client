package main

import "github.com/jason-shen/redis-client/pkg/client"

func main()  {
	client.NewRedisCache("localhost:6379", "", 1, 10)

}