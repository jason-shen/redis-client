# redis client for go
redis client library is to make easy to do crud with json object in redis

## Features
- [x] add json objects
- [x] edit json objects
- [x] remove json object
- [x] add list
- [x] edit list
- [x] remove list items by key
- [x] delete list by key

## Installing the Library
```textmate
git clone https://github.com/jason-shen/redis-client.git
```

## Usage 
**example of SET/GET**
```go
func main()  {
	redis := client.NewRedisCache("localhost:6379", "", 1, 1000)
	key := "myfirst_key"
	object := map[string]interface{} {
		"firstname": "hello",
		"lastname": "world",
	}

	redis.Set(key, object)
	data, err := redis.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", data)
}
```
please see the example folder

## Dependencies
- [go-redis](github.com/go-redis/redis/)

## License
MIT License - see [LICENSE](LICENSE) for full text


