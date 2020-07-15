package dao

import (
	"github.com/go-redis/redis"
)

type redisDao struct {
	rdb *redis.Client
}

// NewRedisDao creates a new Redis Client
func NewRedisDao(c *redis.Client) RedisDao {
	return &redisDao{
		rdb: c,
	}
}

func (r *redisDao) Set(key, value string) error {
	err := r.rdb.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisDao) Get(key string) (value string, err error) {
	v, err := r.rdb.Get(key).Result()
	if err != nil {
		return "", err
	}

	return v, nil
}

func (r *redisDao) Flush() error {
	return r.rdb.FlushAll().Err()
}
