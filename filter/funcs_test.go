package filter

import (
	"github.com/indigo-web/iter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnique(t *testing.T) {
	data := []string{"hello", "world", "hello", "alex", "world"}
	want := []string{"hello", "world", "alex"}
	it := iter.Filter(Unique(make([]string, 0, 5)), iter.Slice(data))

	for i := 0; ; i++ {
		str, cont := it.Next()
		if !cont {
			require.Equal(t, 3, i)
			break
		}

		require.Less(t, i, len(want))
		require.Equal(t, want[i], str)
	}
}
