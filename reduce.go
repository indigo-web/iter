package iter

type reduceFunc[T any] func(prev, curr T) T

func Reduce[T any](fun reduceFunc[T], iter Iterator[T], initial ...T) (acc T) {
	values := Extract(iter, initial)

	if len(values) == 0 {
		return acc
	}

	acc = values[0]

	for _, v := range values[1:] {
		acc = fun(acc, v)
	}

	return acc
}
