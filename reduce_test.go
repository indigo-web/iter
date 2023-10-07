package iter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type nopIter[T any] struct {
}

func (n nopIter[T]) Next() (el T, cont bool) {
	return el, false
}

func (n nopIter[T]) Stopped() bool {
	return false
}

func (n nopIter[T]) Break() {
}

func TestReduce(t *testing.T) {
	t.Run("no initial, no values", func(t *testing.T) {
		res := Reduce[int](nil, nopIter[int]{})
		require.Equal(t, 0, res)
	})

	t.Run("no initial", func(t *testing.T) {
		sum := Reduce[int](func(prev, curr int) int {
			return prev + curr
		}, Slice[int]([]int{1, 2, 3}))

		require.Equal(t, 6, sum)
	})

	t.Run("with initial", func(t *testing.T) {
		sum := Reduce[int](func(prev, curr int) int {
			return prev + curr
		}, Slice[int]([]int{1, 2, 3}), 1)

		require.Equal(t, 7, sum)
	})
}
