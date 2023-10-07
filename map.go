package iter

type MapFunc[A, B any] func(el A) B

type mapIterator[A, B any] struct {
	iter    Iterator[A]
	fun     MapFunc[A, B]
	stopped bool
}

func Map[A, B any](mapFunc MapFunc[A, B], iter Iterator[A]) Iterator[B] {
	return mapIterator[A, B]{
		iter: iter,
		fun:  mapFunc,
	}
}

func (m mapIterator[A, B]) Next() (el B, cont bool) {
	a, cont := m.iter.Next()
	if !cont {
		m.Break()

		return el, false
	}

	return m.fun(a), true
}

func (m mapIterator[A, B]) Stopped() bool {
	return m.stopped
}

func (m mapIterator[A, B]) Break() {
	m.stopped = true
}
