package storage

type Storage interface {
	Name() string
	Get(key string) string
	Set(key string, val string) error
	Delete(key string)
	Update(key string, val string)
}