package client

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

type redisCache struct {
	host string
	password string
	db int
	expires time.Duration
	data Items
}

func NewItems(key string) *Items {
	var item = &Items{
		key: key,
		items: Item{data: make(map[string]interface{})},
	}
	return item
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

func (r redisCache) Set(key string, value string) error {
	client := r.getClient()
	err := client.Set(key, value, r.expires*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r redisCache) Get(key string) (string, error) {
	client := r.getClient()
	// fmt.Println(key)
	val, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisCache) Create(key string, value interface{}) error {
	client := r.getClient()
	//fmt.Println("creating", value)
	item, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Json", item)
	err = client.Set(key, item, r.expires*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Read(key string) (interface{}, error) {
	var value map[string]interface{}
	client := r.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return err, nil
	}

	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		return err, nil
	}
	return value, nil
}

func (r *redisCache) Update(key string, value interface{}) error {
	prev, err := r.Read(key)
	if err != nil {
		return err
	}
	result := prev.(map[string]interface{})

	val := Items{
		key: key,
		items: Item{data: result},
	}

	//fmt.Println("values", val.items.data)
	i := strconv.Itoa(len(val.items.data))
	m := NewItems(key)
	m.items.data[i] = val.items.data


	if val.key == key {
		//fmt.Println("before data", m.items.data, "key", val.key)
		//fmt.Println("input value", value)
		i := strconv.Itoa(len(m.items.data))
		m.items.data[i] = value.(map[string]interface{})
		//fmt.Println("after data", m.items.data, "key", val.key)
		err := r.Create(key, m.items.data)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *redisCache) Delete(key string) error {
	client := r.getClient()
	err := client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}
