package delivery

type Broker[T any] interface {
	Publish(data T)
	Close()
}