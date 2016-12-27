package redislock

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/redsync.v1"
)

var (
	redsyncClient *redsync.Redsync
)

func LoadLockRedis(key string) (err error) {
	hosts := []string{"127.0.0.1:6379"} 
	poolsize:= 3
	timeoutS := 3
	if timeoutS == 0 {
		timeoutS = 60
	}
	timeout := time.Second * time.Duration(timeoutS)
	if poolsize == 0 {
		poolsize = 3
	}
	fmt.Println("redis hosts:", hosts, "timeout seconds:", timeout)
	var pools []redsync.Pool
	for _, host := range hosts {
		pool := &redis.Pool{
			MaxIdle:     poolsize,
			IdleTimeout: time.Second * 240,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.DialTimeout("tcp", host, timeout, timeout, timeout)
				if err != nil {
					return nil, err
				}
				return conn, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
		pools = append(pools, pool)
	}
	redsyncClient = redsync.New(pools)
	return
}

func NewRedisLock(name string) *RedisLock {
	lock := redsyncClient.NewMutex(name, redsync.SetTries(1))
	return &RedisLock{name: name, lock: lock}
}

type RedisLock struct {
	name   string
	lock   *redsync.Mutex
	expire int
}

func (this *RedisLock) Lock() error {
	if this == nil {
		return errors.New("empty of *RedisLock")
	}

	if this.name == "" {
		return errors.New("empty of RedisLock.name")
	}

	if this.lock == nil {
		return errors.New("empty of RedisLock.lock")
	}

	err := this.lock.Lock()

	// keep lock alive
	go func() {
		pingTicker := time.NewTicker(time.Second * 4)
		defer func() {
			recover()
			pingTicker.Stop()
		}()
		for _ = range pingTicker.C {
			if !this.lock.Extend() {
				return
			}
		}
	}()

	return err
}

func (this *RedisLock) Unlock() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if this == nil {
		return errors.New("empty of *RedisLock")
	}

	if this.name == "" {
		return errors.New("empty of RedisLock.name")
	}

	if this.lock == nil {
		return errors.New("empty of RedisLock.lock")
	}

	ok := this.lock.Unlock()
	if ok {
		return nil
	}

	return errors.New("RedisLock unlock failed")
}

//-----------------------------------------------------------------------------
type RedisLockEx struct {
	name   string
	lock   *redsync.Mutex
	expire int
}

func NewRedisLockEx(name string, e int) *RedisLockEx {
	lock := redsyncClient.NewMutex(name, redsync.SetTries(1))
	return &RedisLockEx{name: name, lock: lock, expire: e}
}

func (this *RedisLockEx) LockOnce() error {
	if this == nil {
		return errors.New("empty of *RedisLock")
	}

	if this.name == "" {
		return errors.New("empty of RedisLock.name")
	}

	if this.lock == nil {
		return errors.New("empty of RedisLock.lock")
	}

	err := this.lock.Lock()
	if err != nil {
		return err
	}

	go func() {
		pingTicker := time.NewTicker(time.Second * 4)
		expireT := time.After(time.Second * time.Duration(this.expire))
		defer func() {
			recover()
			pingTicker.Stop()
		}()

		//keep lock alive within the expiration time
		for {
			select {
			case <-pingTicker.C:
				if !this.lock.Extend() {
					return
				}
			case <- expireT:
				return
			}
		}

	}()
	return nil
}
