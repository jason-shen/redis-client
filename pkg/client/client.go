package client

type Redisclient interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Remove(key string)
	Addlist(key string, value interface{}) error
	Getlist(key string, start int64, offset int64) ([]string, error)
	RemoveList(key string, count int64, value interface{}) error
	DeleteList(key string) error
}