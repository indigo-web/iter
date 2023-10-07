package iter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("pass all", func(t *testing.T) {
		iter := Filter(
			func(el string) bool {
				return el != "don't want"
			},
			Slice([]string{"i", "don't want", "want", "pizza"}),
		)
		want := []string{"i", "want", "pizza"}

		for i := 0; ; i++ {
			str, cont := iter.Next()
			if !cont {
				require.Equal(t, 3, i)
				break
			}

			require.Less(t, i, len(want))
			require.Equal(t, want[i], str)
		}
	})

	t.Run("fail all", func(t *testing.T) {
		iter := Filter(
			func(el string) bool {
				return false
			},
			Slice([]string{"i", "don't want", "want", "pizza"}),
		)

		_, cont := iter.Next()
		require.False(t, cont)
	})
}
