package iter

func Extract[T any](from Iterator[T], to []T) []T {
	for {
		el, cont := from.Next()
		if !cont {
			return to
		}

		to = append(to, el)
	}
}
