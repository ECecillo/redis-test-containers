package store

type Store interface {
	GetCounterValue(key string) int
}
