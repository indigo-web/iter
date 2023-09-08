package filter

import "github.com/indigo-web/iter"

func Unique[T comparable](buff []T) iter.FilterFunc[T] {
	return func(el T) bool {
		for _, seen := range buff {
			if el == seen {
				return false
			}
		}

		buff = append(buff, el)

		return true
	}
}
