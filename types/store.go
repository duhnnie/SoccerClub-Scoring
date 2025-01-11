package types

type Store[T any] interface {
	Get(string) (T, bool)
}
