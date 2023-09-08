package iter

type Iterator[T any] interface {
	Next() (el T, cont bool)
	Stopped() bool
	Break()
}
