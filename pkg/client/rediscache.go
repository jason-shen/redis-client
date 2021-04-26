package client

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

type redisCache struct {
	host string
	password string
	db int
	expires time.Duration
}

func NewRedisCache(host string, password string, db int, exp time.Duration) Redisclient {
	return &redisCache{
		host: host,
		password: password,
		db: db,
		expires: exp,
	}
}

func (r *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: r.host,
		Password: r.password,
		DB: r.db,
	})
}

func (r *redisCache) Addlist(key string, value interface{}) error {
	client := r.getClient()
	err := client.LPushX(key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisCache) Getlist(key string, start int64, offset int64) ([]string, error) {
	client := r.getClient()
	value, err := client.LRange(key, start, offset).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r *redisCache) RemoveList(key string, count int64, value interface{}) error {
	client := r.getClient()

	err := client.LRem(key, count, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisCache) DeleteList(key string) error {
	client := r.getClient()

	err := client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r redisCache) Set(key string, value interface{}) error {
	client := r.getClient()
	// fmt.Println(key, value)
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = client.Set(key, json, r.expires*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r redisCache) Get(key string) (interface{}, error) {
	var value interface{}
	client := r.getClient()
	// fmt.Println(key)
	val, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(val), &value)

	return value, nil
	// room :=
}

func (r *redisCache) Remove(key string) {
	client := r.getClient()
	client.Del(key)
}
