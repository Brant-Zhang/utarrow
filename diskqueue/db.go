package diskqueue

type DB interface {
	Close() error
	Write(data []byte) error
}
