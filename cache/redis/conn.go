package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	pool *redis.Pool
	host = "192.168.88.131:6379"
	pass = ""
)

func init() {
	pool = newRedisPool()
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 100,
		MaxActive: 30,
		IdleTimeout: 5*time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			if _, err := c.Do("AUTH", pass); err !=nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisPool() *redis.Pool {
	return pool
}