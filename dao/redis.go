package dao

import (
	"github.com/go-redis/redis"
	"github.com/leewei05/image-api/rest"
)

type redisDao struct {
	client *redis.Client
}

// NewRedisDao creates a new Redis Client
func NewRedisDao(c *redis.Client) rest.RedisDao {
	return &redisDao{
		client: c,
	}
}
