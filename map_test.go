package iter

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	data := []string{"123", "456", "789"}
	want := []int{123, 456, 789}
	iter := Map[string, int](Slice(data), func(el string) int {
		num, _ := strconv.Atoi(el)
		return num
	})

	for i := 0; ; i++ {
		num, cont := iter.Next()
		if !cont {
			require.Equal(t, 3, i)
			break
		}

		require.Less(t, i, len(want))
		require.Equal(t, want[i], num)
	}
}
