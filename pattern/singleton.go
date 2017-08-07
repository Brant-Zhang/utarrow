package pattern

import "sync"

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func NewSingle() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
