package activeLists

import (
	"github.com/go-redis/redis"
	"strings"
	"time"
)

type redisStorage struct {
	cli *redis.Client
}

func (r *redisStorage) Put(list, key string, data map[string]interface{}, ttl time.Duration) error {
	finalKey := r.makeFinalKey(list, key)
	pipe := r.cli.TxPipeline()
	pipe.HMSet(finalKey, data)
	if ttl > 0 {
		pipe.Expire(finalKey, ttl)
	}
	_, err := pipe.Exec()
	return err
}

func (r *redisStorage) Del(list, key string) error {
	return r.cli.Del(r.makeFinalKey(list, key)).Err()
}

func (r *redisStorage) Get(list, key, field string) ([]byte, error) {
	return r.cli.HGet(r.makeFinalKey(list, key), field).Bytes()
}

func (r *redisStorage) makeFinalKey(list, key string) string {
	sb := strings.Builder{}
	sb.WriteString(alNamePrefix)
	sb.WriteString(list)
	sb.WriteByte(':')
	sb.WriteString(key)
	return sb.String()
}

func newRedisStorage(url string) (*redisStorage, error) {
	cli := redis.NewClient(&redis.Options{Addr: url})
	if err := cli.Ping().Err(); err != nil {
		return nil, err
	}
	return &redisStorage{cli: cli}, nil
}
