package redis

// CacheInterface - interface for cache
type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration int) error
}
