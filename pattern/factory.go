package pattern

import (
	"io"
	"os"

	"github.com/go-redis/redis"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type diskStore struct {
	dirPath string
}

func (d *diskStore) Open(file string) (io.ReadWriteCloser, error) {
	return os.OpenFile(d.dirPath+file, os.O_RDWR|os.O_CREATE, 0755)
}

func newDiskStorage() Store {
	return &diskStore{dirPath: "/home/brant/temp/"}
}

type memStore struct {
	r *redis.Client
	k string
}

func newMemoryStorage() Store {
	return new(memStore)
}

func (this *memStore) Open(key string) (io.ReadWriteCloser, error) {
	var err error
	this.r = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	this.k = key
	return this, err
}

func (this *memStore) Read(p []byte) (n int, err error) {
	var v []byte
	v, err = this.r.Get(this.k).Bytes()
	if err != nil {
		return
	}
	n = copy(p, v)
	return
}

func (this *memStore) Write(p []byte) (n int, err error) {
	err = this.r.Set(this.k, string(p), 0).Err()
	if err == nil {
		n = len(p)
	}
	return
}

func (this *memStore) Close() error {
	return this.r.Close()
}

func newTempStorage() (s Store) {
	return s
}

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()
	}
}
