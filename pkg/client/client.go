package client

type Redisclient interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Create(key string, value interface{}) error
	Read(key string) (interface{}, error)
	Update(key string, value interface{}) error
	Delete(key string) error
	Addlist(key string, value interface{}) error
	Getlist(key string, start int64, offset int64) ([]string, error)
	RemoveList(key string, count int64, value interface{}) error
	DeleteList(key string) error
}

type Item struct {
	data map[string]interface{}
}

type Items struct {
	key string
	items Item
}