package redis

import (
	"time"
)

var Client *RedisManager

func Init(host string) (err error) {
	var poolsize int = 5
	var timeout int = 2000
	Client, err = NewRedisManager(
		host, poolsize, time.Millisecond*time.Duration(timeout))
	return
}
