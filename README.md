# redis client for go
redis client library is to make easy to do crud with json object in redis

## Features
- [x] create json objects
- [x] update json objects
- [x] delete json object
- [x] add list
- [x] edit list
- [x] remove list items by key
- [x] delete list by key

## Installing the Library
```textmate
git clone https://github.com/jason-shen/redis-client.git
```

## Usage 
**example of Create/Update**
```go
func main()  {
redis := client.NewRedisCache("localhost:6379", "", 1, 1000)
key := "myfirst_key"
sample := map[string]interface{} {
"firstname": "hello1",
"lastname": "world0",
}

sample2 := map[string]interface{} {
"firstname": "hello2",
"lastname": "world1",
}
sample3 := map[string]interface{} {
"firstname": "hello3",
"lastname": "world2",
}

redis.Create(key, sample)
data, err := redis.Read(key)
if err != nil {
panic(err)
}
fmt.Println("object create", data)
redis.Update(key, sample2)
data2, err := redis.Read(key)
fmt.Println("object update 1", data2)
redis.Update(key, sample3)
data3, err := redis.Read(key)
fmt.Println("object update 2", data3)
}
```
please see the example folder

## Dependencies
- [go-redis](github.com/go-redis/redis/)

## License
MIT License - see [LICENSE](LICENSE) for full text


