package iter

type FilterFunc[T any] func(el T) bool

type filter[T any] struct {
	iter    Iterator[T]
	fun     FilterFunc[T]
	stopped bool
}

func Filter[T any](fun FilterFunc[T], iter Iterator[T]) Iterator[T] {
	return filter[T]{
		iter: iter,
		fun:  fun,
	}
}

func (f filter[T]) Next() (el T, cont bool) {
	for {
		el, cont = f.iter.Next()
		if !cont {
			f.Break()

			return el, false
		}

		if f.fun(el) {
			return el, true
		}
	}
}

func (f filter[T]) Stopped() bool {
	return f.stopped
}

func (f filter[T]) Break() {
	f.stopped = true
}
