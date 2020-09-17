package diskqueue

type DB interface {
	Close() error
	Write(data []byte) error
	Delete() error
	Empty() error
	ReadChan() <-chan []byte
}
