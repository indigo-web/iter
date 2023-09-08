package iter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("iterate till end", func(t *testing.T) {
		data := []string{"i", "want", "pizza"}
		iter := Slice(data)

		for i := 0; ; i++ {
			str, cont := iter.Next()
			if !cont {
				require.Equal(t, 3, i)
				break
			}

			require.Less(t, i, len(data))
			require.Equal(t, data[i], str)
		}
	})

	t.Run("with break", func(t *testing.T) {
		data := []string{"i", "want", "pizza"}
		iter := Slice(data)

		for i := 0; ; i++ {
			_, cont := iter.Next()
			if !cont {
				require.Equal(t, 1, i)
				break
			}

			iter.Break()
		}
	})
}

func TestPairedSlice(t *testing.T) {
	t.Run("iterate till end", func(t *testing.T) {
		data := []string{"i", "want", "pizza", "pepperoni"}
		want := [][]string{
			{"i", "want"}, {"pizza", "pepperoni"},
		}
		iter := PairedSlice(data)

		for i := 0; ; i++ {
			pair, cont := iter.Next()
			if !cont {
				require.Equal(t, 2, i)
				break
			}

			require.Less(t, i, len(want))
			require.Equal(t, want[i], pair)
		}
	})

	t.Run("iterate by even amount of elements", func(t *testing.T) {
		data := []string{"i", "want", "pizza", "pepperoni", "special"}
		want := [][]string{
			{"i", "want"}, {"pizza", "pepperoni"},
		}
		iter := PairedSlice(data)

		for i := 0; ; i++ {
			pair, cont := iter.Next()
			if !cont {
				require.Equal(t, 2, i)
				break
			}

			require.Less(t, i, len(want))
			require.Equal(t, want[i], pair)
		}
	})

	t.Run("with break", func(t *testing.T) {
		data := []string{"i", "want", "pizza", "pepperoni"}
		iter := PairedSlice(data)

		for i := 0; ; i++ {
			_, cont := iter.Next()
			if !cont {
				require.Equal(t, 1, i)
				break
			}

			iter.Break()
		}
	})
}
