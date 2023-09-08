package iter

type sliceIterator[T any] struct {
	slice []T
	ptr   int
}

func Slice[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{
		slice: slice,
	}
}

func (s *sliceIterator[T]) Next() (el T, cont bool) {
	if s.Stopped() {
		return el, false
	}

	el = s.slice[s.ptr]
	s.ptr++

	return el, true
}

func (s *sliceIterator[T]) Stopped() bool {
	return s.ptr >= len(s.slice)
}

func (s *sliceIterator[T]) Break() {
	s.ptr = len(s.slice)
}

type pairedSliceIterator[T any] struct {
	slice   []T
	ptr     int
	stopped bool
}

func PairedSlice[T any](slice []T) Iterator[[]T] {
	return &pairedSliceIterator[T]{
		slice: slice,
	}
}

func (p *pairedSliceIterator[T]) Next() (el []T, cont bool) {
	if p.Stopped() || p.ptr+1 >= len(p.slice) {
		return el, false
	}

	el = p.slice[p.ptr : p.ptr+2]
	p.ptr += 2

	return el, true
}

func (p *pairedSliceIterator[T]) Stopped() bool {
	return p.ptr+1 >= len(p.slice)
}

func (p *pairedSliceIterator[T]) Break() {
	p.ptr = len(p.slice)
}
