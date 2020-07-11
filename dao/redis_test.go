package dao

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

var (
	redisHost = "localhost"
	redisPort = 6379

	testRedisData = []redisData{
		{
			key:   "dog",
			value: "gs://image-api-v1/dog.png",
		},
		{
			key:   "cat",
			value: "gs://image-api-v1/cat.png",
		},
		{
			key:   "rabbit",
			value: "gs://image-api-v1/rabbit.png",
		},
	}
)

type redisData struct {
	key    string
	value  string
	expire time.Duration
}

type redisTestSuite struct {
	suite.Suite
	impl *redisDao
	rdb  *redis.Client
}

func TestRedisDao(t *testing.T) {
	suite.Run(t, new(redisTestSuite))
}

func (r *redisTestSuite) SetupTest() {
	redisStr := fmt.Sprintf("%s:%v", redisHost, redisPort)
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     redisStr,
		Password: "",
		DB:       0,
	})

	_, err := r.rdb.Ping().Result()
	if err != nil {
		log.Panic("Cannot open test Redis")
	}

	newRedis := NewRedisDao(r.rdb)

	r.impl = newRedis.(*redisDao)
}

func (r *redisTestSuite) TestSetAndGet() {
	for _, t := range testRedisData {
		err := r.impl.Set(t.key, t.value)
		r.NoError(err)
	}

	v, err := r.impl.Get("dog")
	r.NoError(err)
	r.Equal("gs://image-api-v1/dog.png", v)
}

func (r *redisTestSuite) TestFlush() {
	err := r.impl.Flush()
	r.NoError(err)

	v, _ := r.impl.Get("dog")
	r.Equal("", v)
}
